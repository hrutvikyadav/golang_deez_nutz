package iteration

import "strings"

// Repeat takes in a string and repeats it 5 times
func Repeat(char string) (repeated string ) {
	var sb strings.Builder
	repeatCount := 5
	for range repeatCount {
		sb.WriteString(char)
	}

	repeated = sb.String()
	return
}

