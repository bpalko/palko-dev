// Package content holds Brendan's portfolio copy, mirrored from apps/web.
package content

import (
	"fmt"
	"strings"
)

const (
	Name     = "Brendan Palkowski"
	Site     = "https://palko.dev"
	Role     = "systems engineer — platform & infra"
	Location = "Grand Rapids, MI"
	Age      = "28"
	School   = "University of Michigan"
	Degree   = "Bachelor of Science in Information Science"
	GitHub   = "https://github.com/bpalko"
	LinkedIn = "https://www.linkedin.com/in/brendan-palkowski/"
	X        = "https://x.com/b_palkowski"
	Resume   = "https://palko.dev/resume"
)

const Intro = "I design the systems other teams ship on: Kubernetes APIs, GitOps, multi-cloud infrastructure. Currently at First American."

const About = `This is a passion site. I'm looking to dive into blogging and writing.

I'm not the world's most well-read person, but the books I have read I'm passionate about. You'll find those on the bookshelf, both here and at home.

You'll find my bar cart. I'm interested in spirits, mainly whiskey and red wines. Recently been experimenting with different gins as well.

You'll also find music, pulled straight from Spotify. No curation, no highlight reel — I promise I have taste.`

type SkillGroup struct {
	Label string
	Items string
}

var Skills = []SkillGroup{
	{Label: "Core", Items: "Go, Kubernetes (operators, CRDs, Kubebuilder), Terraform, Helm, Crossplane, ArgoCD, AWS/GCP"},
	{Label: "Infrastructure", Items: "Postgres, Python, Istio, Kyverno, Docker, GitLab CI/CD, GitHub Actions"},
	{Label: "Agentic", Items: "Cursor (self-hosted agents), Claude Code"},
}

type Job struct {
	Role     string
	Company  string
	Period   string
	Location string
	Bullets  []string
}

var Experience = []Job{
	{
		Role:     "Senior Software Engineer",
		Company:  "First American",
		Period:   "Jan. 2024 – Present",
		Location: "Remote",
		Bullets: []string{
			"Cut PR database wait from ~20 minutes (full Aurora via Crossplane per PR) by building a Go operator that tenants each preview onto a shared Aurora cluster across ~100s concurrent PRs/day.",
			"Authored Crossplane XRDs, compositions, and Helm packages for Aurora RDS/DocDB, S3, Redis, and IAM (Pod Identity), and led adoption into the Golden Path for ~400 application tenants, replacing ticket-driven Terraform.",
			"Built a policy-driven Go operator that provisions and reconciles NetworkPolicies, Pod Security, ResourceQuotas, and related controls in labeled customer namespaces, drift-corrected for the life of each tenant.",
			"Led ArgoCD as the platform's first real CI/CD path, designing Application Set onboarding and sequential environment promotion (dev → stage → prod) so teams could go PR-to-prod on a per-commit cadence instead of multi-month cycles. Every tenant depends on a Helm chart we maintain that abstracts platform features onto Kubernetes, including the Istio mesh, secrets management, autoscaling, and other Kubernetes primitives.",
			"Technical Pod Lead for two years across platform and controller initiatives, accountable for teams of 3-8 while remaining hands-on in operator architecture and implementation; platform on-call.",
		},
	},
	{
		Role:     "Cloud Engineer",
		Company:  "Alteryx",
		Period:   "Mar. 2022 – Jan. 2024",
		Location: "Remote",
		Bullets: []string{
			"Rebuilt the multi-cloud Terraform platform (GCP and AWS) across four accounts plus customer data planes, cutting environment builds from weeks to under 3 hours.",
			"Wrote Go services for split-plane cluster lifecycle, Terraform templating, and GitLab mirroring.",
			"Led AI data-plane infrastructure in Terraform; onboarded teams onto Kubernetes, ArgoCD, and Helm.",
		},
	},
	{
		Role:     "Predictive Modeler",
		Company:  "Auto-Owners Insurance",
		Period:   "June 2020 – Mar. 2022",
		Location: "Lansing, MI",
		Bullets: []string{
			"Built and ran 70+ production scrapers with Kafka and cron, used by underwriters for risk analysis.",
			"Migrated the department to Konvoy Kubernetes (production and non-production) and trained modelers.",
		},
	},
}

type Volunteer struct {
	Role     string
	Org      string
	URL      string
	Location string
	Bullets  []string
}

var PiAcademy = Volunteer{
	Role:     "Coding Instructor",
	Org:      "PiAcademy",
	URL:      "https://thepiacademy.org",
	Location: "Grand Rapids, MI",
	Bullets: []string{
		"Teach coding fundamentals through Raspberry Pi projects at a local youth tech nonprofit.",
	},
}

type Book struct {
	Title  string
	Author string
	Year   int
	Status string // reading | completed | want-to-read
	Notes  string
}

var Books = []Book{
	{Title: "Kitchen Confidential", Author: "Anthony Bourdain", Year: 2000, Status: "completed", Notes: "My favorite writer of all time, if I'm being honest. Everything he writes is visceral and raw. I'm always drawn into his world and perspective on life."},
	{Title: "Halo: The Fall of Reach", Author: "Eric Nylund", Year: 2001, Status: "completed", Notes: "I love the Halo universe; I grew up in it. This is the first book in the \"Halo Series\". Highly recommend reading this to gain some lore on the Master Chief and how he came to be."},
	{Title: "Halo: The Flood", Author: "William C. Dietz", Year: 2003, Status: "completed", Notes: "Continuation of the Halo books, albeit the only one based around the first game, Halo: Combat Evolved. This has been a treat to read as I sunk many hours into the video game."},
	{Title: "The Intelligent Investor", Author: "Benjamin Graham", Year: 1949, Status: "completed", Notes: "One of my first introductions into financial literacy. A dense read. I probably grasped half of the subjects... probably should read it once more."},
	{Title: "The Courage to Be Disliked", Author: "Ichiro Kishimi and Fumitake Koga", Year: 2018, Status: "reading", Notes: "Reading this one off and on. Adlerian psychology and Stoic themes. Compelling."},
	{Title: "The Bobiverse", Author: "Dennis E. Taylor", Year: 2016, Status: "completed", Notes: "A series about a programmer who copies himself across the universe. Please — just read them..."},
	{Title: "As A Man Thinketh", Author: "James Allen", Year: 1903, Status: "completed", Notes: "One of those books you don't just read and forget about; you'll want to read this over and over. I sat and stared at each page for extended periods trying to grasp what Allen is getting across. The ideas have endured the test of time. Give it a shot."},
	{Title: "Going Infinite: The Rise and Fall of a New Tycoon", Author: "Michael Lewis", Year: 2023, Status: "completed", Notes: "I've always had an interest in finances and crypto. SBF is an incredibly interesting person, no matter how you view the events of FTX or his life. Lewis made reading this seem like a movie."},
	{Title: "Meditations", Author: "Marcus Aurelius, translated by Gregory Hays", Year: 2003, Status: "completed", Notes: "It's Meditations. Read it a 100 times. I like the application of the Stoic mindset."},
	{Title: "The Phoenix Project", Author: "Gene Kim, Kevin Behr, George Spafford", Year: 2014, Status: "completed", Notes: "What software engineer hasn't read this? Well they should..."},
	{Title: "The Martian", Author: "Andy Weir", Year: 2011, Status: "completed", Notes: "I've seen the movie many times; I heard the book makes Mark Watney's struggles even more intense."},
	{Title: "Dungeon Crawler Carl", Author: "Matt Dinniman", Year: 2020, Status: "reading", Notes: "A LitRPG where a man and his ex-girlfriend's cat survive a deadly alien game show built inside a dungeon. Ridiculous premise, hard to put down."},
	{Title: "Red Rising", Author: "Pierce Brown", Year: 2014, Status: "reading", Notes: "Sci-fi caste system, rebellion, and a lot of betrayal. Everyone told me to read this one for years."},
	{Title: "The Big Short", Author: "Michael Lewis", Year: 2010, Status: "want-to-read", Notes: "I've seen the movie many times; need to read the book. I've enjoyed Lewis' writing style in the past."},
	{Title: "Medium Raw: A Bloody Valentine to the World of Food and People Who Cook", Author: "Anthony Bourdain", Year: 2010, Status: "want-to-read", Notes: "A follow-up to Kitchen Confidential. I'm on a mission to read all of his books. I'll probably start back up here."},
}

type Bottle struct {
	Name       string
	Type       string
	Distillery string
	Proof      float64
	Rating     int
	Notes      string
	Status     string // current | finished | wishlist
}

var BarCart = []Bottle{
	{Name: "Buffalo Trace", Type: "Bourbon", Distillery: "Buffalo Trace Distillery", Proof: 90, Rating: 4, Notes: "A solid everyday bourbon. You'll be hard pressed to find something better for the price.", Status: "current"},
	{Name: "Midwinter Night's Dram - Act 13", Type: "Whiskey", Distillery: "Highwest Distillery", Proof: 99, Rating: 5, Notes: "I like this year. Warm, vanilla and spices.", Status: "current"},
	{Name: "Midwinter Night's Dram - Act 12", Type: "Whiskey", Distillery: "Highwest Distillery", Proof: 98.6, Rating: 3, Notes: "Kinda average compared to Act 11 and Act 13.", Status: "current"},
	{Name: "EH Taylor Small Batch", Type: "Bourbon", Distillery: "Buffalo Trace Distillery", Proof: 100, Rating: 4, Notes: "It's a solid bourbon. Picked it up while in Kentucky visiting the Buffalo Trace distillery", Status: "current"},
	{Name: "Eagle Rare 10 Year", Type: "Bourbon", Distillery: "Buffalo Trace Distillery", Proof: 90, Rating: 4, Notes: "It's a solid bourbon. Can't go wrong with Eagle Rare. I have a few bottles at this rate.", Status: "current"},
	{Name: "Elijah Craig Barrel Proof - Store Pick Austin TX", Type: "Bourbon", Distillery: "Heaven Hill Distillery", Proof: 129, Rating: 3, Notes: "Picked it up on the way out of Austin TX. A bit too hot for my taste, but decent flavor underneath. The 129 isn't really the issue; I have higher proofs that drink easier.", Status: "current"},
	{Name: "Tenjaku", Type: "Gin", Distillery: "Tenjaku", Proof: 86, Rating: 4, Notes: "Been enjoying this gin for a while now. Smooth with nice botanical flavors. Great for cocktails or sipping neat.", Status: "current"},
	{Name: "Campari", Type: "Aperitif", Distillery: "Campari Group", Proof: 48, Rating: 4, Notes: "Negroni. That is all.", Status: "current"},
	{Name: "Rosso", Type: "Vermouth", Distillery: "Martini & Rossi", Proof: 30, Rating: 4, Notes: "Negroni. Cheap and available sweet vermouth.", Status: "current"},
}

type Playlist struct {
	Name string
	ID   string
}

const OnRepeatID = "37i9dQZF1EpvH5fg4WpI6V"

var Playlists = []Playlist{
	{Name: "Yo Daddy's Yacht Rock", ID: "0D0ExDmtDIOc4rsUgptXde"},
	{Name: "Divorced Dad Rock", ID: "1pQ3pMBkbzxgKsp9H4UxVZ"},
	{Name: "MidWest Emo", ID: "6uPSXhw5JB7i6XvMI5FhBH"},
	{Name: "Illmatic", ID: "19TBXgTlTflStysaFjv1UQ"},
	{Name: "Two Tequila Sodas at ZARA", ID: "7JTKtuTe9spXawzBWjZqfM"},
	{Name: "Driftin To Yo Momma's House", ID: "2yzu2WOeqMZrMEuKcATe9j"},
	{Name: "Palko's Jazz Encyclopedia", ID: "2ereachssECetKrBWAYWHw"},
	{Name: "Pure Testosterone Fuel", ID: "1QD6Itgb3SR5yM0srNciI9"},
	{Name: "GOTHBOICLIQUE", ID: "5pWtwPvf7aXeCoLRkjq4YC"},
}

func SpotifyURL(id string) string {
	return "https://open.spotify.com/playlist/" + id
}

type Post struct {
	Title       string
	Description string
	Date        string
	Tags        []string
	Body        string
}

var Posts = []Post{
	{
		Title:       "Come Back Later",
		Description: "I plan to try my hand at blog writing. Ain't there yet...",
		Date:        "2024-01-20",
		Tags:        []string{"learning"},
		Body:        "Insert Blog Post Here",
	},
}

func BooksByStatus(status string) []Book {
	out := make([]Book, 0)
	for _, b := range Books {
		if b.Status == status {
			out = append(out, b)
		}
	}
	return out
}

func BottlesByStatus(status string) []Bottle {
	out := make([]Bottle, 0)
	for _, b := range BarCart {
		if b.Status == status {
			out = append(out, b)
		}
	}
	return out
}

func Stars(n int) string {
	if n < 0 {
		n = 0
	}
	if n > 5 {
		n = 5
	}
	return strings.Repeat("★", n) + strings.Repeat("☆", 5-n)
}

func FormatProof(p float64) string {
	if p == float64(int(p)) {
		return fmt.Sprintf("%d Proof", int(p))
	}
	s := strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.1f", p), "0"), ".")
	return s + " Proof"
}
