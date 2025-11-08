# palko.dev

A personal portfolio built with Astro and TypeScript, featuring a homepage with resume, blog, and bookshelf.

This is just a passion site for me. I'm looking to dive into the world of blogging and writing. I've drawn a lot of inspiritation from Anthony Bourdain, even if his voice and style may not mesh with technical writing. But I don't want this page to only be technical; we'll see what I turn it in to.

I'm not the world's most well-read person, but the books I have read I'm passionate about. You'll find those on my bookshelf, both in this site, and at home.

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
├── package.json
├── public/           # Static assets (images, favicon, etc.)
├── README.md
├── astro.config.mjs
├── tsconfig.json
└── src/
	├── components/   # Reusable components
	├── data/         # TypeScript data files
	│   └── books.ts  # Bookshelf data
	├── layouts/      # Page layouts
	│   ├── BaseLayout.astro
	│   └── BlogPost.astro
	├── pages/        # Routes (file-based routing)
	│   ├── index.astro           # Homepage
	│   ├── bookshelf.astro       # Bookshelf page
	│   └── blog/
	│       ├── index.astro       # Blog listing
	│       └── *.md              # Blog posts
	└── utils/        # Various utils in TypeScript (API lookups)
```
