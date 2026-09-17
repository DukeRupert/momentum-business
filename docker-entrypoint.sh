#!/bin/sh
set -e

# No logging from this script: busybox `date` has no nanosecond support, so it
# cannot emit the standard's timestamp format, and anything non-JSON written
# here would break the single-format rule. The API and Caddy both log their own
# startup on stdout.

# Start the Go API in the background
if [ -f /usr/local/bin/contact-api ]; then
    /usr/local/bin/contact-api &
    sleep 1
fi

# Start Caddy in the foreground
exec caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
