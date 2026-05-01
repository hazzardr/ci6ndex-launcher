# Civ 6 Mod Manager — Future Work

This document tracks features, enhancements, and architectural improvements to consider after the MVP is stable and shipped.

---

## Features

### Steam Workshop Automation
- **Auto-subscribe**: Use `steam://openurl/<workshop_url>` or SteamCMD integration to automatically subscribe to missing Workshop mods when importing a profile.
- **Workshop browser**: Search and browse Steam Workshop directly in the app using the `IPublishedFileService` API.
- **Mod update notifications**: Detect when subscribed Workshop mods have updated since the profile was last applied.

### Save Game Integration
- **Save compatibility warnings**: Parse Civ 6 save files to detect which mods were active when the save was created. Warn users if their current profile does not match.
- **Save-to-profile**: Extract the mod list from an existing save and generate a new profile from it.

### Profile Management
- **Profile categories/tags**: Organize profiles into folders (e.g., "Multiplayer", "Single Player", "Total Conversion").
- **Profile versioning**: Keep a history of changes to a profile so users can roll back.
- **Profile dependencies**: Allow profiles to inherit from or extend other profiles.
- **Load order visualization**: Display and manually adjust mod load order within a profile.
- **Profile notes/description**: Rich text descriptions per profile.

### UI / UX
- **Dark/light theme**: Follow system theme or manual toggle.
- **Mod thumbnails**: Fetch and cache Workshop preview images.
- **Drag-and-drop mod organizer**: Reorder and assign mods to profiles via drag-and-drop.
- **Tray icon / quick switch**: Minimize to system tray; right-click to switch profiles without opening the full app.
- **Onboarding wizard**: First-run setup to auto-detect Civ 6 and Steam paths.

### Safety & Reliability
- **Automatic backups**: Scheduled backups of `Mods.sqlite` with retention policy.
- **Profile dry-run**: Preview what changes will be made to `Mods.sqlite` before applying.
- **Conflict detection**: Detect mods known to conflict with each other and warn users.
- **Civ 6 process guard**: Block profile switching while Civ 6 is running; optionally auto-close and relaunch.

### Platforms
- **macOS & Linux support**: Detect paths for macOS (`~/Library/...`) and Linux (`~/.local/share/...` / Proton prefixes).
- **Microsoft Store / Xbox Game Pass**: Support alternate install paths for non-Steam versions of Civ 6.

### Backend & Sharing
- **Profile expiration / deletion**: Allow users to delete or set TTL on shared profiles.
- **User accounts (optional)**: Claim ownership of shared profiles, list all profiles created by a user.
- **Profile popularity / discovery**: Public directory of popular community profiles.
- **Import from URL**: Support importing profiles from a direct JSON URL or file.

### CLI & Automation
- **CLI mode**: Expose all desktop functionality as a CLI for power users and scripting.
- **CI/CD integration**: GitHub Action or script to validate a profile JSON against installed mods.
- **Auto-launch with Civ 6**: Hook into Steam launch options to apply a default profile before starting the game.

### Performance
- **Parallel mod scanning**: Concurrent parsing of `.modinfo` files.
- **SQLite WAL mode**: Use Write-Ahead Logging for safer/faster `Mods.sqlite` access.
- **Incremental scans**: Only re-scan changed mod directories.

---

## Technical Debt / Refactoring

- **Migrate to Wails v3**: When Wails v3 reaches stable release, evaluate migration for improved performance and multi-window support.
- **Frontend framework migration**: Re-evaluate Svelte vs. Astro vs. Solid if Wails adds first-class Astro support.
- **Shared package expansion**: Move more logic (e.g., `.modinfo` parsing, Steam API types) into `pkg/shared` if a CLI or third-party tools need them.
- **Testing**: Add Go unit tests for `modscan`, `sqlite`, and `profile` packages. Add frontend component tests (Vitest + Testing Library).
- **i18n**: Extract all user-facing strings for localization.

---

## Stretch Goals

- **Total conversion support**: Handle total overhaul mods that replace core game files.
- **Mod diffing**: Compare two profiles to see added/removed/changed mods.
- **Cloud sync**: Sync profiles to a user’s cloud storage (Dropbox, Google Drive, etc.) as an alternative to the share-code backend.
- **Discord Rich Presence**: Show active profile name and mod count in Discord status.
