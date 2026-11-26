import { defineConfig } from 'astro/config';

// Write the static site to the repo-root dist/ folder so the existing
// Cloudflare Pages project (root directory `/`, output `dist`) keeps working
// after the monorepo move.
export default defineConfig({
  site: 'https://palko.dev',
  outDir: '../../dist',
});
