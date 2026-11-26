# SSH portfolio

A terminal UI for [palko.dev](https://palko.dev). Built with the Charm stack from [Joe Sluis's guide](https://www.joesluis.com/blog/ssh-portfolio-guide) (`wish`, `bubbletea`, `bubbles`, `lipgloss`).

```bash
ssh ssh.palko.dev
```

Locally:

```bash
ssh -p 2222 localhost
```

Copy is pulled from the Astro site in `apps/web` (about, resume, bookshelf, bar cart, music, blog, contact). Theme is the same black / green as the website.

## Quick start

```bash
cd apps/ssh
go test ./...
go run ./cmd/ssh-portfolio
```

In another terminal:

```bash
ssh -o StrictHostKeyChecking=no -p 2222 localhost
```

Wish creates `.ssh/host_ed25519` on first run if it is missing. That path is gitignored.

### Make targets

| Target | What it does |
| --- | --- |
| `make test` | `go test ./...` |
| `make run` | `go run ./cmd/ssh-portfolio` |
| `make build` | binary at `bin/ssh-portfolio` |
| `make keys` | generate a host key |
| `make build-linux` | linux/amd64 binary |

### Keys

| Key | Action |
| --- | --- |
| `1`–`7` | About, Experience, Bookshelf, Bar Cart, Music, Blog, Contact |
| `tab` / `h` `l` / arrows | switch sections |
| `j` `k` / arrows | scroll |
| `g` / `G` | top / end |
| `q` / `ctrl+c` | quit |

Resize the terminal: the layout listens for `tea.WindowSizeMsg` and compact-renders below ~56 columns.

## Config

| Env | Default |
| --- | --- |
| `SSH_ADDR` | `0.0.0.0:2222` |
| `HOST_KEY_PATH` | `.ssh/host_ed25519` |

No password or key is required. The server accepts SSH `none` auth so visitors land in the TUI immediately.

## Hosting

The website stays on Cloudflare Pages. This process cannot run there (no raw TCP/22). **Fly.io + a Cloudflare DNS-only record for `ssh.palko.dev` is the recommended path.** Spectrum, Tunnel, and Containers are documented with tradeoffs in [HOSTING.md](./HOSTING.md).
