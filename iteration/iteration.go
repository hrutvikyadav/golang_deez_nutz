package iteration

import "strings"

// Repeat takes in a string and repeats it count times
func Repeat(char string, count int) (repeated string ) {
	var sb strings.Builder
	repeatCount := count
	for range repeatCount {
		sb.WriteString(char)
	}

	repeated = sb.String()
	return
}

