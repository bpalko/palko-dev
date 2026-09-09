# palko.dev

A personal site built with Astro and TypeScript. Home is a short pitch. The resume lives on its own page, with a PDF download.

This is just a passion site for me. I'm looking to dive into the world of blogging and writing. I've drawn a lot of inspiritation from Anthony Bourdain, even if his voice and style may not mesh with technical writing. But I don't want this page to only be technical; we'll see what I turn it in to.

I'm not the world's most well-read person, but the books I have read I'm passionate about. You'll find those on my bookshelf, both in this site, and at home.

You'll find my bar cart. I'm interested in spirits, mainly whiskey and red wines. Recently been experimenting with different gins as well.

You'll also find my Music page, pulled straight from Spotify. No curation, no highlight reel — I promise I have taste.

## 🚀 Quick Start

### Development

```bash
# Install dependencies
npm install --ignore-scripts

# Start development server
npm run dev
```

Visit http://localhost:4321 to see the site!

### Build for Production

```bash
# Build the site
npm run build

# Preview production build locally
npm run preview
```

## 📁 Project Structure

```
/
├── README.md
├── astro.config.mjs
├── package-lock.json
├── package.json
├── public
│   ├── bottles
│   │   ├── buffalotrace.png
│   │   ├── campari.png
│   │   ├── eaglerare.png
│   │   ├── ehtaylorSB.png
│   │   ├── elijahcraigBP.png
│   │   ├── midwinter12.png
│   │   ├── midwinter13.png
│   │   ├── rosso.png
│   │   └── tenjaku.png
│   ├── palko.png
│   └── resume.pdf
├── src
│   ├── components
│   │   └── Resume.astro
│   ├── data
│   │   ├── barCart.ts
│   │   └── books.ts
│   ├── layouts
│   │   ├── BaseLayout.astro
│   │   └── BlogPost.astro
│   ├── pages
│   │   ├── bar-cart.astro
│   │   ├── blog
│   │   │   ├── index.astro
│   │   │   └── learning-typescript.md
│   │   ├── bookshelf.astro
│   │   ├── index.astro
│   │   ├── music.astro
│   │   └── resume.astro
│   └── utils
│       └── bookCovers.ts
└── tsconfig.json
```
