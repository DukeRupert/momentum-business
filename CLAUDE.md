# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Brochure business site for Momentum Business Solutions, LLC (momentumbusiness.org) - a bookkeeping and financial consulting company. Built with Hugo static site generator and a Go API backend for contact form handling.

## Commands

```bash
# Development (requires npm install first)
npm install
hugo server             # Start dev server on localhost:1313

# Build
hugo --gc --minify      # Build for production

# Go API (in api/ directory)
cd api && go run .      # Start API server on localhost:8080

# Docker (production)
docker compose up -d    # Run production container
docker compose logs     # View logs
```

## Architecture

### Tech Stack
- **Static Site**: Hugo with Go templates
- **Styling**: Tailwind CSS v3 (`tailwind.config.js` scans `layouts/` and `content/`) with custom typography system
- **Interactivity**: Alpine.js (loaded via CDN in `baseof.html`)
- **API**: Go 1.21 HTTP server (stdlib only, no external deps) for contact form
- **Deployment**: Docker (multi-stage build) with Caddy reverse proxy
- **Email**: Postmark for transactional emails
- **Bot Protection**: Cloudflare Turnstile (site key in `hugo.toml`, secret key in env)
- **Analytics**: Plausible (self-hosted at plausible.angmar.dev)

### Key Directories
- `api/` - Go API backend (main.go, handlers.go, validation.go, email.go)
- `assets/css/main.css` - Tailwind base + custom typography system
- `content/` - Hugo markdown content pages
- `data/` - Structured YAML content (services, FAQ, navigation, team, testimonials)
- `layouts/` - Hugo Go templates (`_default/baseof.html` is the base, `partials/` has reusable components)
- `static/` - Static assets (images, favicons)

### Custom Typography System

Defined in `assets/css/main.css` with a newspaper-inspired type scale:
- Custom font classes: `.text-kicker`, `.text-caption`, `.text-body`, `.text-subhead`, `.text-headline-*`, `.text-display-*`
- Single font family: `--font-primary` and `--font-display` both use Inter
- Color palette: `--primary-*` (cloud-burst navy, based on #1C2657)
- Custom color utility classes (`.bg-primary-*`, `.text-primary-*`, etc.) since these are CSS custom properties not in the Tailwind config

### Contact Form Flow

1. User submits form on `/contact` (Alpine.js handles client-side)
2. Honeypot field (`website`) silently rejects bots with a fake success response
3. Cloudflare Turnstile token verified server-side
4. POST to `/api/contact` with Go validation (regex patterns for name/email/phone, enum checks for revenue/services)
5. On success: sends notification email to business owner + thank-you email to user via Postmark
6. Client-side redirects to `/success?name=...&email=...`

### Docker Architecture

Single container runs both Caddy and the Go API via `docker-entrypoint.sh`:
- Go API starts as a background process on port 8080
- Caddy runs in foreground, reverse-proxies `/api/*` to the Go API, serves Hugo static files from `/srv`
- Multi-stage build: Hugo (hugomods/hugo:exts) → Go compile → Caddy final image

### Environment Variables

Required in `.env` for production:
```
POSTMARK_TOKEN=xxx           # Postmark API token
POSTMARK_TO=cade@momentumbusiness.org
POSTMARK_FROM=noreply@momentumbusiness.org
ALLOWED_ORIGINS=https://www.momentumbusiness.org  # Comma-separated, defaults to http://localhost:1313
TURNSTILE_SECRET_KEY=xxx     # Cloudflare Turnstile secret (skips verification if unset)
API_PORT=8080                # Go API port (defaults to 8080, use to avoid conflicts)
```

### Deployment

GitHub Actions (`.github/workflows/deploy.yml`): push to `main` → Docker build → push to Docker Hub → SSH deploy to VPS at `/opt/momentum-business`.
