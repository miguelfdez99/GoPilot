package main

import (
	"fmt"
	"io/ioutil"
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

func getLinuxDistribution() (string, error) {
	fileContent, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(fileContent))
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID_LIKE" {
			return strings.Trim(parts[1], "\""), nil
		}
	}

	// If ID_LIKE is not present, try ID(Fedora does not have ID_LIKE)
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			continue
		}
		if strings.Trim(parts[0], "\"") == "ID" {
			return strings.Trim(parts[1], "\""), nil
		}
	}

	return "", fmt.Errorf("unable to determine Linux distribution")
}
