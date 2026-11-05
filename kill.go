package main

import (
	"fmt"
	"strings"
)

type Kill struct {
	Time     string
	Killer   string
	Victim   string
	Weapon   string
	Headshot bool
}

func ParseKill(line string) (Kill, error) {
	fields := strings.Fields(line)
	if len(fields) < 6 || len(fields) > 7 {
		return Kill{}, fmt.Errorf("invalid kill line: %q", line)
	}

	if len(fields) == 7 && fields[6] != "headshot" {
		return Kill{}, fmt.Errorf("invalid kill line: %q", line)
	}

	if fields[2] != "killed" {
		return Kill{}, fmt.Errorf("invalid kill line: %q", line)
	}
	if fields[4] != "with" {
		return Kill{}, fmt.Errorf("invalid kill line: %q", line)
	}
	kill := Kill{Time: fields[0], Killer: fields[1], Victim: fields[3], Weapon: fields[5]}
	if len(fields) == 7 && fields[6] == "headshot" {
		kill.Headshot = true
	}

	return kill, nil

}
