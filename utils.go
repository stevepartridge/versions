package versions

import (
	"strings"
)

func Parse(ver string) (string, string, string) {
	parts := strings.Split(ver, ".")
	if len(parts) == 0 {
		return "", "", ""
	} else if len(parts) == 1 {
		return parts[0], "", ""
	} else if len(parts) == 2 {
		return parts[0], parts[1], ""
	} else if len(parts) == 3 {
		return parts[0], parts[1], parts[2]
	} else if len(parts) > 3 {
		return parts[0], parts[1], strings.Join(parts[2:], "-")
	}
	return "", "", ""
}
