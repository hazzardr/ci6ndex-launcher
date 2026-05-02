# README

## About

Wails Svelte + TypeScript + less + Prettier + ESLint + Vite Template.

The template is based on the official Wails svelte-ts template, upgraded to Svelte 5, manually installed less, Prettier and ESLint. The configuration files are from official SvelteKit template created using `npx sv create`.

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend` directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in yourbrowser and connect to your application.

## Code Checking, Linting and Formatting

To check the code, run `npm run check` in the `frontend` directory. Or run `npm run check:watch` to watch for changes.

To lint the code, run `npm run lint` in the `frontend` directory.

To format the code, run `npm run format` in the `frontend` directory.

## Building

To build a redistributable, production mode package, use `wails build`.
