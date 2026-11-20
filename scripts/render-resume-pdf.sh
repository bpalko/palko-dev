#!/usr/bin/env bash
# Regenerate public/resume.pdf from the built resume page.
# Requires: npm run build first, plus Microsoft Edge on the Windows host.
set -euo pipefail
cd "$(dirname "$0")/.."

node scripts/build-resume-print.mjs

WIN_TEMP="/mnt/c/Users/Brendan/AppData/Local/Temp"
EDGE="/mnt/c/Program Files (x86)/Microsoft/Edge/Application/msedge.exe"
HTML_SRC="public/resume-print.html"
HTML_WIN="$WIN_TEMP/resume-print.html"
PDF_WIN="$WIN_TEMP/resume.pdf"

cp "$HTML_SRC" "$HTML_WIN"
rm -f "$PDF_WIN"

"$EDGE" --headless=new --disable-gpu --no-pdf-header-footer \
  --print-to-pdf="C:\\Users\\Brendan\\AppData\\Local\\Temp\\resume.pdf" \
  "file:///C:/Users/Brendan/AppData/Local/Temp/resume-print.html"

# Edge returns before the file flush sometimes
for _ in $(seq 1 40); do
  [[ -s "$PDF_WIN" ]] && break
  sleep 0.25
done

cp "$PDF_WIN" public/resume.pdf
rm -f "$HTML_SRC" "$HTML_WIN" "$PDF_WIN"
ls -la public/resume.pdf
