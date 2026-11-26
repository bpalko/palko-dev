# palko.dev

A monorepo for Brendan Palkowski's site and SSH portfolio.

| App | Path | Hosting |
| --- | --- | --- |
| Website | [`apps/web`](apps/web) | **Cloudflare Pages** at [palko.dev](https://palko.dev) |
| SSH TUI | [`apps/ssh`](apps/ssh) | Fly.io (or any VM) + Cloudflare **DNS-only** for `ssh.palko.dev` — [why](apps/ssh/HOSTING.md) |

Home is a short pitch. The resume lives on its own page, with a PDF download. The same copy shows up over SSH.

This is a passion site. I'm looking to dive into blogging and writing. I've drawn a lot of inspiration from Anthony Bourdain, even if his voice and style may not mesh with technical writing. I don't want this page to only be technical.

You'll find a bookshelf, a bar cart, and a Music page pulled straight from Spotify.

```bash
# website
npm install --ignore-scripts
npm run dev          # http://localhost:4321

# SSH portfolio
cd apps/ssh
go run ./cmd/ssh-portfolio
# then: ssh -p 2222 localhost
```

## Website (Cloudflare)

Astro + TypeScript. Cloudflare Pages should keep using the **repo root**:

- **Build command:** `npm run build`
- **Output directory:** `dist`

`apps/web` writes into that root `dist/` folder so the existing Pages project does not need a new root directory.

```bash
npm run build     # astro check && astro build
npm run preview
```

## SSH portfolio

Charm stack (`wish`, `bubbletea`, `bubbles`, `lipgloss`). See [`apps/ssh/README.md`](apps/ssh/README.md) for keys and local run, and [`apps/ssh/HOSTING.md`](apps/ssh/HOSTING.md) for Cloudflare vs Fly.io.

Cloudflare Pages / Workers cannot listen on TCP 22. The site stays on Cloudflare; `ssh.palko.dev` should be a grey-cloud DNS record at a small SSH origin.

## Layout

```
/
├── apps/
│   ├── web/                 # Astro site (Cloudflare Pages)
│   │   ├── astro.config.mjs
│   │   ├── package.json
│   │   ├── public/
│   │   └── src/
│   └── ssh/                 # Wish SSH TUI
│       ├── cmd/ssh-portfolio/
│       ├── internal/
│       ├── Dockerfile
│       └── fly.toml
├── package.json             # npm workspaces → @palko/web
└── README.md
```
