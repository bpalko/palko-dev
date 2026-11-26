package tui

import (
	"strings"
	"unicode"
)

func wrapText(s string, width int) string {
	if width < 8 {
		width = 8
	}
	var out []string
	for _, para := range strings.Split(s, "\n") {
		if strings.TrimSpace(para) == "" {
			out = append(out, "")
			continue
		}
		out = append(out, wrapLine(para, width)...)
	}
	return strings.Join(out, "\n")
}

func wrapLine(s string, width int) []string {
	words := splitWords(s)
	if len(words) == 0 {
		return []string{""}
	}
	var lines []string
	var cur strings.Builder
	for _, w := range words {
		if cur.Len() == 0 {
			if len(w) <= width {
				cur.WriteString(w)
				continue
			}
			lines = append(lines, chunk(w, width)...)
			continue
		}
		if cur.Len()+1+len(w) <= width {
			cur.WriteByte(' ')
			cur.WriteString(w)
			continue
		}
		lines = append(lines, cur.String())
		cur.Reset()
		if len(w) <= width {
			cur.WriteString(w)
			continue
		}
		parts := chunk(w, width)
		lines = append(lines, parts[:len(parts)-1]...)
		cur.WriteString(parts[len(parts)-1])
	}
	if cur.Len() > 0 {
		lines = append(lines, cur.String())
	}
	return lines
}

func wrapIndent(prefix, s string, width int) string {
	hang := strings.Repeat(" ", len([]rune(prefix)))
	bodyWidth := width - len([]rune(prefix))
	if bodyWidth < 8 {
		bodyWidth = 8
		prefix = ""
		hang = ""
	}
	lines := wrapLine(s, bodyWidth)
	if len(lines) == 0 {
		return prefix
	}
	var b strings.Builder
	b.WriteString(prefix)
	b.WriteString(lines[0])
	for _, line := range lines[1:] {
		b.WriteByte('\n')
		b.WriteString(hang)
		b.WriteString(line)
	}
	return b.String()
}

func splitWords(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return strings.FieldsFunc(s, unicode.IsSpace)
}

func chunk(s string, width int) []string {
	if width < 1 {
		return []string{s}
	}
	var out []string
	r := []rune(s)
	for len(r) > width {
		out = append(out, string(r[:width]))
		r = r[width:]
	}
	if len(r) > 0 {
		out = append(out, string(r))
	}
	return out
}
