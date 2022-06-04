package maker

import "strings"

func cleanLine(line string) string {
	newL := strings.TrimSpace(line)
	if line[len(line)-1] != '.' {
		newL += "."
	}

	return newL
}
