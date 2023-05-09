package tools

import "strings"

func EscapeUUID(uuid string) string {
	return strings.ReplaceAll(uuid, "-", "\\-")
}
