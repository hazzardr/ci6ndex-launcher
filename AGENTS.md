# AGENTS.md — ci6ndex-launcher

## What this repo is

A Civ 6 mod profile manager: a Wails v2 desktop app (Go + Vue 3) that scans installed mods, lets users create/switch mod profiles, and shares them via short codes.

## Monorepo layout

- `go.work` — Go workspace. Modules:
  - `./pkg/shared` — shared Go types (`Mod`, `Profile`)
  - `./apps/launcher` — Wails desktop app (Go backend + Vue 3 frontend)
- `apps/launcher/frontend/` — Vue 3 + TypeScript + Vite + less. **Not Svelte** (the README is stale).
- `examples/mods/` — sample Workshop mod data for reference.

## Build & dev commands

### Prerequisites
- `go` 1.26.2, `wails` v2 CLI, `pnpm`

### Launcher app
```bash
cd apps/launcher

# Dev (requires two terminals)
wails dev                          # terminal 1: starts backend + serves embedded frontend
cd frontend && pnpm run dev        # terminal 2: Vite dev server on localhost:34115

# Build frontend only
cd frontend && pnpm run build      # outputs to frontend/dist; required before `go vet`/`wails build`

# Build redistributable
wails build
wails build -platform windows      # cross-compile

# Frontend checks
cd frontend
pnpm run check          # vue-tsc --noEmit
pnpm run check:watch    # vue-tsc --watch
pnpm run lint           # prettier --check + eslint
pnpm run format         # prettier --write
pnpm run test           # vitest run
```

### Go
```bash
# From repo root or apps/launcher
go vet ./...
go build ./...

# Linting
golangci-lint run      # config at .golangci.yml; very strict (60+ linters enabled)
```

**Note**: `go vet` and `go build` fail with `pattern all:frontend/dist: no matching files found` until you run `pnpm run build` in the frontend directory at least once. The `//go:embed all:frontend/dist` directive in `main.go` requires the dist folder to exist.

## Key architecture notes

- **Wails bindings**: Public methods on `App` (in `app.go`) are auto-exposed to the frontend. Import them from `../wailsjs/go/main/App.js`.
- **Shared types**: `pkg/shared/mod.go` and `profile.go`. Imported via workspace replace: `github.com/hazzardr/ci6ndex-launcher/shared`.
- **Mod discovery** (`internal/modscan/`):
  - Workshop path on Linux: `~/.steam/steam/steamapps/workshop/content/289070/`
  - Parses `.modinfo` XML (extension is `.modinfo`, not `.modinfo.xml`)
- **Frontend auto-imports**: `unplugin-auto-import` makes Vue and VueUse APIs available without imports. `unplugin-vue-components` auto-imports components from `src/components/`. Generated type files (`auto-imports.d.ts`, `components.d.ts`) are gitignored.
- **Path alias**: `~/` maps to `src/` in both Vite and TypeScript.
- **Vue macros**: `propsDestructure` and `defineModel` are enabled via `unplugin-vue-macros`.
- **Tests**: Vitest + jsdom for frontend unit tests. `src/composables/*.test.ts` is a valid pattern.
- **No Go tests yet** — there are zero `*_test.go` files in the repo.
- **No CI/CD** — `.github/workflows/` and `scripts/` do not exist.

## Style & conventions

- Go: `golangci-lint` config is extremely strict (`cyclop`, `funlen`, `mnd`, `nakedret`, etc.). `package-comments` is disabled.
- Frontend: Prettier + ESLint flat config (`eslint.config.ts`). Vue recommended rules.
- Imports in Go use the workspace replace path for shared types.

## Gotchas

1. **README lies about the frontend framework**: `apps/launcher/README.md` says Svelte; the real stack is **Vue 3** (see `package.json`, `vite.config.ts`, `App.vue`).
2. **Frontend build required before Go build**: `main.go` embeds `frontend/dist`. Run `pnpm run build` first or `wails build` handles it.
3. **Linux Workshop path**: `steamModFolderLocation()` checks `~/.steam/steam/...` first, then falls back to `XDG_DATA_HOME` / `~/.local/share/Steam/...`.
4. **Wails dev mode**: Must run `pnpm run dev` separately in `frontend/` for HMR. The `wails dev` command does not auto-start the Vite dev server in this setup.
