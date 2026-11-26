# Hosting the SSH portfolio

The website (`apps/web`) stays on **Cloudflare Pages** at [palko.dev](https://palko.dev). This Go process is a raw SSH server. Cloudflare cannot replace the origin for `ssh ssh.palko.dev` the way Pages replaces a static web host.

## What Cloudflare can and cannot do

| Product | Raw `ssh host` on TCP/22? | Notes |
| --- | --- | --- |
| **Pages / Workers** | No | HTTP, HTTPS, and WebSockets only. They cannot listen on TCP 22. |
| **Workers Static Assets** | No | Same HTTP-only surface as Pages. |
| **Cloudflare Containers** | Not as a public SSH port | `wrangler containers ssh` is **admin debug**, authenticated to your Cloudflare account. It does **not** publish port 22. Inbound TCP to a Container is a **public beta** (Nov 2026): Spectrum sends TCP to a Worker `connect()` handler, which can pipe bytes to a container port. That needs a **Pro/Business** zone (one Worker-origin Spectrum app per zone, one TCP port) or Spectrum Enterprise. Free zones are out. |
| **Spectrum** | Yes, as a proxy | Pro/Business include **one SSH app per zone** (`tcp/22`). You still run this binary on an origin. Spectrum sits in front for DDoS / anycast. Paid add-on on Pro/Business. |
| **Tunnel (`cloudflared`)** | Not for visitors | Origin needs no inbound firewall hole. Clients must use `ProxyCommand cloudflared access ssh --hostname %h`, **or** a browser-rendered terminal behind Cloudflare Access. That is not `ssh ssh.palko.dev`. |
| **DNS-only (grey cloud)** | Yes | An A/AAAA record with the proxy **off** points `ssh.palko.dev` straight at a VM/Fly machine running this binary. `palko.dev` itself stays orange-clouded for the website. |

Orange-clouding an SSH hostname is a common footgun: Cloudflare's HTTP proxy only speaks the documented HTTP(S) ports. SSH on 22 will not pass through a proxied record unless Spectrum is in front.

## Recommended path

Keep the site on Cloudflare. Run this SSH server on **Fly.io** (or any small VM). Use Cloudflare only as **authoritative DNS**.

1. Deploy `apps/ssh` with the Dockerfile / `fly.toml` in this directory.
2. In Cloudflare DNS, add a **DNS-only** record:

   | Type | Name | Content | Proxy |
   | --- | --- | --- | --- |
   | A (or AAAA) | `ssh` | Fly/VM public IP | DNS only (grey cloud) |

3. Visitors run:

   ```bash
   ssh ssh.palko.dev
   ```

4. Leave `palko.dev` and `www` proxied (orange cloud) so the Astro site stays on Pages.

Why Fly (or a VM) instead of forcing Cloudflare: Pages already hosts the website; Spectrum is extra cost for a hobby TUI; Tunnel breaks the `ssh host` UX; Containers + Spectrum Worker-origin is beta and still needs a paid plan. A tiny always-on machine on port 22 is the same pattern Joe Sluis's guide uses, and it fits this repo.

### Fly.io sketch

```bash
cd apps/ssh
fly apps create palko-ssh
fly volumes create ssh_host_keys --region ord --size 1
fly deploy
fly ips list   # put this IPv4/IPv6 on the grey-cloud ssh record
```

The volume keeps the Wish host key stable so clients do not get a new TOFU warning on every deploy.

Local Docker (if you have it):

```bash
docker build -t palko-ssh .
docker run --rm -p 2222:2222 -v palko-ssh-keys:/data palko-ssh
ssh -p 2222 localhost
```

## Optional: Spectrum in front

If the zone is Pro/Business and you want Cloudflare anycast + DDoS on SSH:

1. Keep the origin running (Fly/VM).
2. Create a Spectrum **SSH** application for `ssh.palko.dev` → `tcp://ORIGIN_IP:22`.
3. See `spectrum-app.example.json`.

You still need the origin. Spectrum does not host the process.

## Optional: Tunnel (no inbound 22)

Use `cloudflared.yml.example`. Tell visitors they need:

```sshconfig
Host ssh.palko.dev
  ProxyCommand cloudflared access ssh --hostname %h
```

or a browser terminal at `https://ssh.palko.dev` with Access + browser rendering. This is a worse public-portfolio UX than grey-cloud DNS.

## What not to do

- Do not add a root `wrangler.toml` that tries to bind TCP 22 on Workers. It will not work, and it can confuse the existing Pages project.
- Do not orange-cloud `ssh.palko.dev` without Spectrum.
- Do not reuse Cloudflare Containers' `authorized_keys` / `wrangler containers ssh` as the public portfolio. That is operator access, not `ssh ssh.palko.dev`.
