package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result := []map[string]any{}

	for _, d := range data {
		fields := strings.Split(d, ";")
		person := map[string]any{}

		// Parsing name
		person["name"] = fields[0]

		// Parsing age
		if age, err := strconv.Atoi(fields[1]); err == nil {
			person["age"] = age
		}

		// Parsing address
		person["address"] = fields[2]

		// Parsing height, if exists
		if fields[3] != "" {
			if height, err := strconv.ParseFloat(fields[3], 64); err == nil {
				person["height"] = height
			}
		}

		// Parsing is_married, if exists
		if len(fields) > 4 && fields[4] != "" {
			if isMarried, err := strconv.ParseBool(fields[4]); err == nil {
				person["isMarried"] = isMarried
			}
		}

		result = append(result, person)
	}

	return result
}
