package main

import (
	"cmp"
	"slices"
	"strings"
)

type PlayerStats struct {
	Name      string
	Kills     int
	Deaths    int
	Headshots int
}

func BuildStats(kills []Kill) map[string]*PlayerStats {
	playerStats := make(map[string]*PlayerStats)
	for _, kill := range kills {
		killer := getOrCreate(playerStats, kill.Killer)
		killer.Kills++
		if kill.Headshot {
			killer.Headshots++
		}

		victim := getOrCreate(playerStats, kill.Victim)
		victim.Deaths++
	}

	return playerStats
}

func (p PlayerStats) KD() float64 {
	if p.Deaths == 0 {
		return float64(p.Kills)
	}
	return float64(p.Kills) / float64(p.Deaths)
}

func (p PlayerStats) HSPercent() float64 {
	if p.Kills == 0 {
		return 0
	}
	return float64(p.Headshots) / float64(p.Kills) * 100
}

func Leaderboard(stats map[string]*PlayerStats) []PlayerStats {
	leaders := make([]PlayerStats, 0, len(stats))
	for _, p := range stats {
		leaders = append(leaders, *p)
	}
	slices.SortFunc(leaders, func(a, b PlayerStats) int {
		if a.Kills != b.Kills {
			return cmp.Compare(b.Kills, a.Kills)
		}
		return strings.Compare(a.Name, b.Name)
	})
	return leaders
}

func getOrCreate(stats map[string]*PlayerStats, name string) *PlayerStats {
	pStats, ok := stats[name]
	if !ok {
		pStats = &PlayerStats{Name: name}
		stats[name] = pStats
	}
	return pStats
}
