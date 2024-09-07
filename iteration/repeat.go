package iteration

func Repeat(char string, count int) string {
	var repeated = ""
	for i := 0; i < count; i++ {
		repeated += char
	}
	return repeated
}
