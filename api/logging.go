package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// contextKey is a private type so our context keys can't collide with others'.
type contextKey string

const requestIDKey contextKey = "request_id"

// trustedProxyHops is how many proxies that append to X-Forwarded-For sit
// between this process and the internet. The chain is:
//
//	visitor -> host Caddy (writes the visitor IP) -> container Caddy
//	           (appends its peer) -> this API
//
// so the right-most entry is the Docker bridge address and the entry we want
// is one to its left. Raise this if another appending proxy (e.g. a second
// reverse proxy) is added; a CDN that Caddy is configured to trust adds its
// own entry to the left and does not change this number.
const trustedProxyHops = 2

// setupLogger installs a JSON slog logger on stdout as the default logger.
// slog's JSONHandler already emits `time`, uppercase `level` and `msg`; we
// only force the timestamp to UTC so every service agrees on a timezone.
func setupLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if len(groups) == 0 && a.Key == slog.TimeKey {
				a.Value = slog.TimeValue(a.Value.Time().UTC())
			}
			return a
		},
	}))
	slog.SetDefault(logger)
	return logger
}

// newRequestID returns a 16-character hex id.
func newRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand failing is not worth dropping the request over; a
		// timestamp-derived id is still unique enough to correlate lines.
		return strconv.FormatInt(time.Now().UnixNano(), 16)
	}
	return hex.EncodeToString(b)
}

// requestIDFrom returns the request id carried on the context, or "".
func requestIDFrom(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return ""
}

// logFor returns a logger with the context's request_id already attached, so
// every line emitted while serving a request can be correlated.
func logFor(ctx context.Context) *slog.Logger {
	if id := requestIDFrom(ctx); id != "" {
		return slog.With("request_id", id)
	}
	return slog.Default()
}

// clientIP returns the real client IP, without a port.
//
// X-Forwarded-For is only honoured because this service is never exposed
// directly to the internet — it listens on localhost and is reached through
// the reverse proxies counted by trustedProxyHops. We take the entry our own
// trusted proxy wrote (counting from the right), never the left-most one,
// which is caller-supplied and forgeable.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		// Index of the entry written by the outermost proxy we trust.
		i := len(parts) - trustedProxyHops
		if i < 0 {
			// Shorter chain than expected (e.g. a hop was reconfigured to
			// replace rather than append) — fall back to the left-most entry
			// we were given rather than reading out of bounds.
			i = 0
		}
		if ip := stripPort(strings.TrimSpace(parts[i])); ip != "" {
			return ip
		}
	}
	return stripPort(r.RemoteAddr)
}

// stripPort removes a trailing :port from an address if present.
func stripPort(addr string) string {
	if addr == "" {
		return ""
	}
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}

// statusRecorder captures the status code and whether anything was written.
type statusRecorder struct {
	http.ResponseWriter
	status int
	err    string
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Write(b []byte) (int, error) {
	if rec.status == 0 {
		rec.status = http.StatusOK
	}
	return rec.ResponseWriter.Write(b)
}

// setError records the error string reported on a failed request.
func (rec *statusRecorder) setError(msg string) {
	rec.err = msg
}

// recorderKey lets handlers attach an error string to the access log line.
const recorderKey contextKey = "recorder"

// noteRequestError attaches an error string to this request's access log line.
// Call it alongside writing a 5xx response.
func noteRequestError(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if rec, ok := ctx.Value(recorderKey).(*statusRecorder); ok {
		rec.setError(err.Error())
	}
}

// loggingMiddleware assigns a request id, puts it on the context and emits
// exactly one access log line per request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		id := newRequestID()
		rec := &statusRecorder{ResponseWriter: w}

		ctx := context.WithValue(r.Context(), requestIDKey, id)
		ctx = context.WithValue(ctx, recorderKey, rec)
		w.Header().Set("X-Request-Id", id)

		next.ServeHTTP(rec, r.WithContext(ctx))

		if rec.status == 0 {
			rec.status = http.StatusOK
		}

		// The health check is polled continuously by the monitor; keep it out
		// of normal traffic at DEBUG.
		level := slog.LevelInfo
		msg := "request"
		if r.URL.Path == "/api/health" {
			level = slog.LevelDebug
		} else if rec.status >= 500 {
			level = slog.LevelError
			msg = "request failed"
		}

		attrs := []any{
			"request_id", id,
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"status", rec.status,
			"duration_ms", float64(time.Since(start).Microseconds()) / 1000.0,
			"remote_ip", clientIP(r),
			"user_agent", r.UserAgent(),
			"referer", r.Referer(),
		}
		if rec.err != "" {
			attrs = append(attrs, "error", rec.err)
		}

		slog.Default().Log(r.Context(), level, msg, attrs...)
	})
}
