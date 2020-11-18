package etl

import "strings"

func Transform(given map[int][]string) map[string]int {

	result := make(map[string]int)

	for key, value := range given {
		for _, s := range value {
			result[strings.ToLower(s)] = key
		}
	}
	return result
}
