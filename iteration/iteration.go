package iteration

import "strings"

func Repeat(ch string, count int) string {
	var sb strings.Builder

	for i := 0; i < count; i++ {
		sb.WriteString(ch)
	}
	return sb.String()
}
