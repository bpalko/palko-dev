import { readFileSync, writeFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";

const root = join(dirname(fileURLToPath(import.meta.url)), "..");
const html = readFileSync(join(root, "dist/resume/index.html"), "utf8");

const styleMatch = html.match(/<style>([\s\S]*?)<\/style>/);
const paperMatch = html.match(/<article class="paper"[\s\S]*?<\/article>/);
if (!styleMatch || !paperMatch) {
  throw new Error("could not extract resume paper/styles from dist/resume/index.html");
}

const out = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>Resume — Brendan Palkowski</title>
<style>
${styleMatch[1]}
html, body { margin: 0; background: #fff; }
.paper { box-shadow: none !important; }
</style>
</head>
<body>
${paperMatch[0]}
</body>
</html>
`;

const outPath = join(root, "public/resume-print.html");
writeFileSync(outPath, out);
console.log(`wrote ${outPath}`);
