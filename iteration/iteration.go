package iteration

import "strings"

func Repeat(ch string, count int) string {
	// var repeated string
	// for i := 0; i < count; i++ {
	// 	repeated += ch
	// }
	// return repeated
	var sb strings.Builder

	for i := 0; i < count; i++ {
		sb.WriteString(ch)
	}
	return sb.String()
}

// func NewRepeat(ch string, count int) string {
// }
