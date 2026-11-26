package content

import (
	"strings"
	"testing"
)

func TestPortfolioHasRealCopy(t *testing.T) {
	if Name != "Brendan Palkowski" {
		t.Fatalf("name: %q", Name)
	}
	if !strings.Contains(Intro, "First American") {
		t.Fatal("intro should mention First American")
	}
	if len(Experience) != 3 {
		t.Fatalf("jobs: got %d", len(Experience))
	}
	if Experience[0].Company != "First American" {
		t.Fatalf("current job: %s", Experience[0].Company)
	}
	if len(Books) < 10 {
		t.Fatalf("books: got %d", len(Books))
	}
	if len(BarCart) != 9 {
		t.Fatalf("bottles: got %d", len(BarCart))
	}
	if len(Playlists) != 9 {
		t.Fatalf("playlists: got %d", len(Playlists))
	}
	if len(Posts) == 0 || Posts[0].Title != "Come Back Later" {
		t.Fatalf("blog: %+v", Posts)
	}
	if PiAcademy.Org != "PiAcademy" {
		t.Fatal("volunteer missing")
	}
}

func TestBooksByStatus(t *testing.T) {
	reading := BooksByStatus("reading")
	if len(reading) != 3 {
		t.Fatalf("reading: %d", len(reading))
	}
	completed := BooksByStatus("completed")
	if len(completed) == 0 {
		t.Fatal("expected completed books")
	}
	want := BooksByStatus("want-to-read")
	if len(want) != 2 {
		t.Fatalf("want-to-read: %d", len(want))
	}
}

func TestStarsAndProof(t *testing.T) {
	if Stars(4) != "★★★★☆" {
		t.Fatalf("stars: %q", Stars(4))
	}
	if Stars(5) != "★★★★★" {
		t.Fatalf("stars 5: %q", Stars(5))
	}
	if FormatProof(90) != "90 Proof" {
		t.Fatalf("proof int: %q", FormatProof(90))
	}
	if FormatProof(98.6) != "98.6 Proof" {
		t.Fatalf("proof frac: %q", FormatProof(98.6))
	}
}
