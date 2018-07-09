package iteration

const repeatCount = 5

// Takes a string and a number of times to repeat that string, and returns that string repeated that number of times
func Repeat(word string, repeat int) string {
	var result string

	if repeat == 0 {
		repeat = repeatCount
	}

	for i := 0; i < repeat; i++ {
		result += word
	}
	return result
}
