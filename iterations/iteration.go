package iterations

import "strings"

func Repeat(character string, time int) string {
	var repeated strings.Builder
	for range time {
		repeated.WriteString(character)
	}
	return repeated.String()
}
