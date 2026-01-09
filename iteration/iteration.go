package iteration

// Repeat takes in a string and repeats it 5 times
func Repeat(char string) (repeated string ) {
	repeatCount := 5
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return
}

