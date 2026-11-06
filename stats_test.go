package main

import "testing"

func TestBuildStats(t *testing.T) {
	kills := []Kill{
		{Time: "20:15:32", Killer: "s1mple", Victim: "NiKo", Weapon: "ak47", Headshot: true},
		{Time: "20:16:01", Killer: "s1mple", Victim: "ZywOo", Weapon: "awp"},
		{Time: "20:16:40", Killer: "NiKo", Victim: "s1mple", Weapon: "m4a1", Headshot: true},
	}

	stats := BuildStats(kills)

	want := map[string]PlayerStats{
		"s1mple": {Name: "s1mple", Kills: 2, Deaths: 1, Headshots: 1},
		"NiKo":   {Name: "NiKo", Kills: 1, Deaths: 1, Headshots: 1},
		"ZywOo":  {Name: "ZywOo", Kills: 0, Deaths: 1, Headshots: 0},
	}
	if len(stats) != len(want) {
		t.Fatalf("got %d players, want %d", len(stats), len(want))
	}
	for name, w := range want {
		got, ok := stats[name]
		if !ok {
			t.Errorf("player %q not found", name)
			continue
		}
		if *got != w {
			t.Errorf("player %q: got %+v, want %+v", name, *got, w)
		}
	}
}

func TestKD(t *testing.T) {
	p := PlayerStats{Kills: 3, Deaths: 2}
	if got := p.KD(); got != 1.5 {
		t.Errorf("KD() = %v, want 1.5", got)
	}

	noDeaths := PlayerStats{Kills: 3, Deaths: 0}
	if got := noDeaths.KD(); got != 3 {
		t.Errorf("KD() with 0 deaths = %v, want 3", got)
	}
}

func TestHSPercent(t *testing.T) {
	p := PlayerStats{Kills: 4, Headshots: 1}
	if got := p.HSPercent(); got != 25 {
		t.Errorf("HSPercent() = %v, want 25", got)
	}

	noKills := PlayerStats{Kills: 0}
	if got := noKills.HSPercent(); got != 0 {
		t.Errorf("HSPercent() with 0 kills = %v, want 0", got)
	}
}

func TestLeaderboard(t *testing.T) {
	stats := map[string]*PlayerStats{
		"ZywOo":  {Name: "ZywOo", Kills: 1},
		"s1mple": {Name: "s1mple", Kills: 2},
		"NiKo":   {Name: "NiKo", Kills: 1},
	}

	got := Leaderboard(stats)

	// more kills first; equal kills -> by name A-Z
	want := []string{"s1mple", "NiKo", "ZywOo"}
	if len(got) != len(want) {
		t.Fatalf("got %d players, want %d", len(got), len(want))
	}
	for i, name := range want {
		if got[i].Name != name {
			t.Errorf("place %d: got %q, want %q", i+1, got[i].Name, name)
		}
	}
}
