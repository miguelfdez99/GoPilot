package main

import (
	"strings"
)

func ExtractFirstParams(input string) []string {
	lines := strings.Split(input, "\n")
	params := make([]string, len(lines))
	for i, line := range lines {
		fields := strings.Split(line, "/")
		params[i] = fields[0]
	}
	return params
}
