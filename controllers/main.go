package controllers

import (
	"strings"
)

func csvToMap(data string) map[string][]string {
	lines := strings.Split(data, "\n")

	columns := strings.Split(lines[0], ",")

	table := make(map[string][]string)
	for _, column := range columns {
		table[column] = make([]string, 0)
	}

	for i := 1; i < len(lines); i++ {
		line := strings.Split(lines[i], ",")
		for j, value := range line {
			table[columns[j]] = append(table[columns[j]], value)
		}
	}
	return table
}
