package tui

import (
	"fmt"
	"strings"

	"github.com/bpalko/palko-dev/ssh/internal/content"
)

func (m Model) sectionContent() string {
	width := m.vp.Width
	if width < 20 {
		width = 20
	}
	switch m.section {
	case sectionAbout:
		return m.viewAbout(width)
	case sectionExperience:
		return m.viewExperience(width)
	case sectionBooks:
		return m.viewBooks(width)
	case sectionBar:
		return m.viewBar(width)
	case sectionMusic:
		return m.viewMusic(width)
	case sectionBlog:
		return m.viewBlog(width)
	case sectionContact:
		return m.viewContact(width)
	default:
		return ""
	}
}

func (m Model) viewAbout(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("About") + "\n\n")
	b.WriteString(st.dim.Render(content.Location+" · "+content.Age) + "\n\n")
	b.WriteString(wrapText(content.Intro, width) + "\n\n")
	b.WriteString(wrapText(content.About, width) + "\n\n")
	b.WriteString(st.heading.Width(width).Render("Education") + "\n\n")
	b.WriteString(st.title.Render(content.School) + "  " + st.maize.Render("GO BLUE") + "\n")
	b.WriteString(st.dim.Render(content.Degree+" · Ann Arbor, MI") + "\n")
	return b.String()
}

func (m Model) viewExperience(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Experience") + "\n\n")
	for _, job := range content.Experience {
		b.WriteString(st.accent.Bold(true).Render(job.Role) + "\n")
		b.WriteString(st.title.Render(job.Company) + st.muted.Render("  "+job.Period+" · "+job.Location) + "\n")
		for _, bullet := range job.Bullets {
			b.WriteString(wrapIndent("  • ", bullet, width) + "\n")
		}
		b.WriteString("\n")
	}

	b.WriteString(st.heading.Width(width).Render("Skills") + "\n\n")
	for _, g := range content.Skills {
		b.WriteString(wrapIndent("  "+g.Label+": ", g.Items, width) + "\n")
	}
	b.WriteString("\n")
	b.WriteString(st.heading.Width(width).Render("Volunteer") + "\n\n")
	v := content.PiAcademy
	b.WriteString(st.accent.Render(v.Role+", "+v.Org) + st.muted.Render("  "+v.Location) + "\n")
	b.WriteString(st.dim.Render("  "+v.URL) + "\n")
	for _, bullet := range v.Bullets {
		b.WriteString(wrapIndent("  • ", bullet, width) + "\n")
	}
	return b.String()
}

func (m Model) viewBooks(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Bookshelf") + "\n")
	b.WriteString(st.dim.Render(wrapText("Books I'm reading, have read, and want to read.", width)) + "\n\n")
	m.writeBooks(&b, width, "Currently Reading", content.BooksByStatus("reading"))
	m.writeBooks(&b, width, "Completed", content.BooksByStatus("completed"))
	m.writeBooks(&b, width, "Want to Read", content.BooksByStatus("want-to-read"))
	return b.String()
}

func (m Model) writeBooks(b *strings.Builder, width int, heading string, books []content.Book) {
	if len(books) == 0 {
		return
	}
	st := m.styles
	b.WriteString(st.accent.Bold(true).Render(heading) + "\n")
	for _, book := range books {
		b.WriteString(st.title.Render(book.Title) + "\n")
		b.WriteString(st.dim.Render(fmt.Sprintf("  %s · %d", book.Author, book.Year)) + "\n")
		if book.Notes != "" {
			b.WriteString(st.muted.Italic(true).Render(wrapIndent("  ", book.Notes, width)) + "\n")
		}
		b.WriteString("\n")
	}
}

func (m Model) viewBar(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Bar Cart") + "\n")
	b.WriteString(st.dim.Render(wrapText("The booze currently on the bar cart. You'll find just about anything.", width)) + "\n\n")
	for _, bottle := range content.BottlesByStatus("current") {
		b.WriteString(st.title.Render(bottle.Name) + "\n")
		meta := fmt.Sprintf("  %s · %s · %s", bottle.Type, content.FormatProof(bottle.Proof), bottle.Distillery)
		b.WriteString(st.accent.Render(bottle.Type) + st.dim.Render(strings.TrimPrefix(meta, "  "+bottle.Type)) + "\n")
		b.WriteString("  " + st.gold.Render(content.Stars(bottle.Rating)) + "\n")
		if bottle.Notes != "" {
			b.WriteString(st.muted.Italic(true).Render(wrapIndent("  ", bottle.Notes, width)) + "\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m Model) viewMusic(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Music") + "\n")
	b.WriteString(st.dim.Render("I promise I have taste...") + "\n\n")
	b.WriteString(st.accent.Bold(true).Render("On Repeat") + "\n")
	b.WriteString("  " + hyperlink("Spotify playlist", content.SpotifyURL(content.OnRepeatID)) + "\n\n")
	b.WriteString(st.accent.Bold(true).Render("Playlists") + "\n")
	for _, p := range content.Playlists {
		url := content.SpotifyURL(p.ID)
		b.WriteString("  " + st.title.Render(p.Name) + "\n")
		b.WriteString("    " + st.dim.Render(hyperlink(url, url)) + "\n")
	}
	_ = width
	return b.String()
}

func (m Model) viewBlog(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Blog") + "\n")
	b.WriteString(st.dim.Render("Thoughts, tutorials, and things I'm learning") + "\n\n")
	for _, post := range content.Posts {
		b.WriteString(st.muted.Render(post.Date) + "\n")
		b.WriteString(st.title.Render(post.Title) + "\n")
		b.WriteString(wrapText(post.Description, width) + "\n\n")
		b.WriteString(wrapText(post.Body, width) + "\n")
		if len(post.Tags) > 0 {
			b.WriteString(st.accent.Render("#"+strings.Join(post.Tags, "  #")) + "\n")
		}
	}
	return b.String()
}

func (m Model) viewContact(width int) string {
	st := m.styles
	var b strings.Builder
	b.WriteString(st.heading.Width(width).Render("Contact") + "\n\n")
	b.WriteString(wrapText("Feel free to reach out. The website is on Cloudflare; this TUI is a sibling app in the same repo.", width) + "\n\n")
	lines := [][2]string{
		{"Web", content.Site},
		{"GitHub", content.GitHub},
		{"LinkedIn", content.LinkedIn},
		{"X", content.X},
		{"Resume", content.Resume},
	}
	for _, line := range lines {
		b.WriteString(st.accent.Render(fmt.Sprintf("  %-9s", line[0])) + hyperlink(line[1], line[1]) + "\n")
	}
	b.WriteString("\n")
	b.WriteString(st.dim.Render("  "+content.Location) + "\n")
	b.WriteString(st.maize.Render("GO BLUE") + "\n")
	return b.String()
}

func hyperlink(label, url string) string {
	return "\x1b]8;;" + url + "\x1b\\" + label + "\x1b]8;;\x1b\\"
}

func stripANSI(s string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == 0x1b {
			if i+1 < len(s) && s[i+1] == '[' {
				j := i + 2
				for j < len(s) && (s[j] < '@' || s[j] > '~') {
					j++
				}
				if j < len(s) {
					i = j + 1
					continue
				}
			}
			if i+1 < len(s) && s[i+1] == ']' {
				if j := strings.Index(s[i:], "\x1b\\"); j != -1 {
					i += j + 2
					continue
				}
			}
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}
