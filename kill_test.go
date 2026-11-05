package main

import "testing"

func TestParseKill(t *testing.T) {
	line := "20:15:32 s1mple killed NiKo with ak47 headshot"

	got, err := ParseKill(line)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := Kill{
		Time:     "20:15:32",
		Killer:   "s1mple",
		Victim:   "NiKo",
		Weapon:   "ak47",
		Headshot: true,
	}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestParseKillWithoutHeadshot(t *testing.T) {
	line := "20:16:01 NiKo killed ZywOo with awp"

	got, err := ParseKill(line)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := Kill{
		Time:     "20:16:01",
		Killer:   "NiKo",
		Victim:   "ZywOo",
		Weapon:   "awp",
		Headshot: false,
	}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestParseKillBadLine(t *testing.T) {
	line := "20:17:00 round started"

	_, err := ParseKill(line)
	if err == nil {
		t.Errorf("expected error for line %q, got nil", line)
	}
}
