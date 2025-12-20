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
- **Styling**: Tailwind CSS v3 with custom typography system
- **API**: Go HTTP server for contact form
- **Deployment**: Docker (multi-stage build) + Caddy
- **Email**: Postmark for transactional emails

### Project Structure
```
├── api/                     # Go API backend
│   ├── main.go              # HTTP server, CORS middleware
│   ├── handlers.go          # Contact form handler
│   ├── validation.go        # Form validation logic
│   └── email.go             # Postmark email templates
├── assets/css/main.css      # Tailwind + custom typography
├── content/                 # Hugo content (markdown)
│   ├── _index.md            # Homepage
│   ├── about.md, services.md, contact.md, privacy.md, success.md
├── data/                    # Structured content
│   ├── services.yaml        # Service packages & pricing
│   ├── faq.yaml             # FAQ content
│   ├── navigation.yaml      # Nav structure
│   └── team.yaml            # Team members
├── layouts/
│   ├── _default/baseof.html # Base template with Alpine.js
│   ├── partials/            # Reusable components
│   │   ├── head.html, header.html, footer.html
│   │   ├── hero.html, features.html, faq.html, cta.html
│   │   ├── cleanup.html, premium.html, consulting.html
│   │   └── mission.html, team.html
│   ├── index.html           # Homepage layout
│   ├── about/, services/, contact/, success/  # Page layouts
├── static/                  # Static assets (images, favicons)
├── hugo.toml                # Hugo configuration
├── Dockerfile               # Multi-stage build
├── Caddyfile                # Caddy reverse proxy config
└── docker-compose.yml       # Container orchestration
```

### Custom Typography System

The site uses a newspaper-inspired type scale defined in `assets/css/main.css`:
- Custom font classes: `.text-kicker`, `.text-caption`, `.text-body`, `.text-subhead`, `.text-headline-*`, `.text-display-*`
- Two font families: `--font-primary` (Manrope) for body, `--font-display` (Outfit) for headings
- Color palette: `--primary-*` (everglade green) and `--secondary-*` (cloud-burst blue)

### Contact Form Flow

1. User submits form on `/contact` (Alpine.js handles client-side)
2. POST to `/api/contact` (Go API)
3. Go validates with custom validation rules matching original Zod schema
4. On success: sends notification to business owner + thank-you email to user via Postmark
5. Client-side redirect to `/success?name=...&email=...`

### Environment Variables

Required in `.env` for production:
```
POSTMARK_TOKEN=xxx           # Postmark API token
POSTMARK_TO=cade@momentumbusiness.org
POSTMARK_FROM=noreply@momentumbusiness.org
ALLOWED_ORIGIN=https://www.momentumbusiness.org
```

### Deployment

The site deploys via GitHub Actions to a VPS:
1. Push to `main` triggers build
2. Docker multi-stage build: Hugo → Go → Caddy
3. Image pushed to Docker Hub
4. SSH deploy pulls and restarts container

See `ARCHITECTURE.md` for detailed deployment documentation.

### Reference Files

The original SvelteKit implementation is preserved in `/sveltekit/` for reference during transition.
