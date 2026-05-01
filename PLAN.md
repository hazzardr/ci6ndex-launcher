# Civ 6 Mod Manager — MVP Plan

## Overview

Build a desktop mod profile manager for Civilization VI, similar to r2modman. Users create named mod profiles (sets of enabled mods), switch between them with one click, and share profiles via short codes. The app manages Civ 6’s `Mods.sqlite` directly so profile switching is instant and safe.

**Tech Stack**
- **Desktop App**: Wails v2 (Go backend + Svelte frontend)
- **Backend Server**: Go 1.26 + Chi router
- **Mono-repo**: Go workspace (`go.work`) with shared packages

> **Note on Astro**: Astro is not available as a Wails frontend template. Wails expects a single-page application with client-side routing. Astro is primarily a static-site generator and would require awkward workarounds. **Svelte** is recommended instead — it is lightweight, has excellent Wails templates, and feels native in a desktop context.

---

## Mono-repo Structure

```
ci6ndex-bundle/
├── go.work                          # Go workspace file
├── PLAN.md                          # This file
├── FUTURE.md                        # Post-MVP roadmap
├── README.md                        # Project overview
├── pkg/
│   └── shared/                      # Shared Go types (desktop ↔ server)
│       ├── profile.go
│       └── mod.go
├── apps/
│   ├── desktop/                     # Wails application
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── wails.json
│   │   ├── frontend/                # Svelte + Vite frontend
│   │   │   ├── src/
│   │   │   ├── package.json
│   │   │   └── vite.config.js
│   │   └── internal/
│   │       ├── config/              # User settings & paths
│   │       ├── modscan/             # Mod discovery (local + Workshop)
│   │       ├── profile/             # Local profile CRUD
│   │       ├── sqlite/              # Mods.sqlite read/write
│   │       ├── steam/               # Steam Web API client
│   │       └── civ6/                # Civ 6 launcher
│   └── server/                      # Containerized Go HTTP API
│       ├── go.mod
│       ├── main.go
│       ├── TODO.md                  # HTTPS, deployment, infra notes
│       └── internal/
│           ├── api/                 # Chi HTTP handlers
│           ├── store/               # Profile persistence
│           └── model/               # Server-side types
└── scripts/
    └── build.sh                     # Cross-platform build helper
```

---

## Phase 1 — Foundation (Week 1)

### 1.1 Repo Bootstrap
- Initialize Go workspace (`go work init`)
- Scaffold `pkg/shared/` with core types:
  - `Mod` (UUID, Name, Version, Author, Source, Path, WorkshopID)
  - `Profile` (ID, Name, Description, ModUUIDs, CreatedAt)
- Initialize Wails project in `apps/desktop/` using the **Svelte** template
- Initialize Go HTTP server in `apps/server/` with Chi router and standard project layout

### 1.2 Desktop — Mod Discovery
Implement `internal/modscan/` to:
- Locate Civ 6 mod directories:
  - Local mods: `Documents/My Games/Sid Meier's Civilization VI/Mods/`
  - Steam Workshop: `<Steam>/steamapps/workshop/content/289070/`
- Parse `.modinfo` XML to extract `id`, `version`, `Name`, `Authors`, etc.
- Resolve Steam Workshop items: map folder names (Workshop IDs) to mod metadata
- Expose a Wails binding: `ScanMods() → []Mod`

### 1.3 Desktop — Profile CRUD (Local Only)
Implement `internal/profile/` to:
- Store profiles as JSON in the app’s config directory
- Methods: `CreateProfile`, `GetProfiles`, `UpdateProfile`, `DeleteProfile`
- Bind to frontend: profile list view, create/edit forms

### 1.4 Desktop — Civ 6 SQLite Integration
Implement `internal/sqlite/` to:
- Open `Mods.sqlite` (read-only for discovery, read-write for switching)
- Map the schema:
  - `mods` table (or equivalent) — identify which columns represent enabled/disabled state
  - Understand how Civ 6 tracks Workshop vs local mods
- Implement `ApplyProfile(profile) → error`:
  - Backup current `Mods.sqlite` before modification
  - Disable all mods
  - Enable only mods present in the profile (by UUID)
- Implement `LaunchCiv6() → error` (optional for MVP, can launch via Steam URI)

**Key Research Task**: Inspect `Mods.sqlite` schema on a machine with Civ 6 installed. Document exact table/column names for mod state.

---

## Phase 2 — Sharing Backend (Week 2)

### 2.1 Server Scaffold
- Chi router with middleware: logging, request ID, CORS, recovery
- Health endpoint: `GET /health`
- Configuration via env vars: `PORT`, `DATABASE_URL`, `API_KEY` (optional rate limiting)

### 2.2 Profile Storage
Implement `POST /profiles`:
- Accepts `Profile` JSON payload
- Generates a short share code (8-character URL-safe base64 / Crockford-style)
- Stores profile JSON in SQLite/Postgres with code as primary key
- Returns `{ "code": "a1b2-c3d4" }`

Implement `GET /profiles/{code}`:
- Retrieves profile JSON by code
- Returns `404` if not found
- Returns `410` if expired (optional for MVP)

### 2.3 Desktop — Share & Import
- **Share**: `ShareProfile(profileID) → string` binding
  - POSTs profile JSON to backend
  - Copies share code to clipboard
  - Displays QR code (nice-to-have)
- **Import**: `ImportProfile(code) → Profile` binding
  - GETs profile from backend
  - Compares mod list against locally installed mods
  - Shows "Missing Mods" dialog with Steam Workshop links (`steam://openurl/...`)
  - Creates local profile copy after user confirms

---

## Phase 3 — Polish & Packaging (Week 3)

### 3.1 Frontend UX
- **Mod Browser**: Two-pane view (available mods on left, profile mods on right)
- **Profile Switcher**: Dropdown or sidebar to select active profile; one-click apply
- **Missing Mods Dialog**: Lists mods not installed, with direct Steam Workshop subscribe links
- **Status Bar**: Shows last applied profile, Civ 6 path, scan status

### 3.2 Error Handling & Safety
- Always backup `Mods.sqlite` before writing
- Validate all mod UUIDs exist before applying profile
- Graceful handling when Civ 6 is running (warn user)
- Clear error messages for path detection failures

### 3.3 Build & Distribution
- Wails build for Windows (`wails build -platform windows`)
- Server container image (`Dockerfile` in `apps/server/`)
- `scripts/build.sh` to automate desktop builds

---

## MVP Success Criteria

- [ ] User can scan all installed Civ 6 mods (local + Workshop)
- [ ] User can create, edit, rename, and delete mod profiles
- [ ] User can apply a profile, which correctly updates `Mods.sqlite`
- [ ] User can share a profile as a short code
- [ ] Another user can import a profile by code and see which mods are missing
- [ ] Desktop app runs on Windows (primary target)

---

## Open Questions / Research Needed

1. **Mods.sqlite schema**: Need exact table/column names for enabled mod state. Must be inspected on a Civ 6 installation.
2. **Steam Workshop API key**: May need a Steam Web API key for reliable mod metadata lookups. Fallback: parse `.modinfo` only.
3. **Civ 6 process detection**: Should the app warn if Civ 6 is running when applying a profile? (Likely yes.)
