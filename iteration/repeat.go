package iteration

func Repeat(char string, loopSize int) string {
  var repeatedAnswer string

	for i := 0; i < loopSize; i++ {
		repeatedAnswer += char
	}

	return repeatedAnswer
}
