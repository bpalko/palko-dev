# Portfolio Website

A personal portfolio built with Astro and TypeScript, featuring a homepage with resume, blog, and bookshelf.

## 🚀 Quick Start

### Development

```bash
# Install dependencies
npm install --ignore-scripts

# Start development server
npm run dev
```

Visit http://localhost:4321 to see your site!

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
├── public/          # Static assets (images, favicon, etc.)
├── src/
│   ├── components/  # Reusable components
│   ├── data/        # TypeScript data files
│   │   └── books.ts # Bookshelf data
│   ├── layouts/     # Page layouts
│   │   ├── BaseLayout.astro
│   │   └── BlogPost.astro
│   └── pages/       # Routes (file-based routing)
│       ├── index.astro           # Homepage
│       ├── bookshelf.astro       # Bookshelf page
│       └── blog/
│           ├── index.astro       # Blog listing
│           └── *.md              # Blog posts
├── astro.config.mjs
├── package.json
└── tsconfig.json
```
