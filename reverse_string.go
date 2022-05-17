package reverse_string

func ReverseString(input string) (output string) {
	tempStr := []rune(input)
	length := len(tempStr)
	for i, j := 0, length-1; i < j; i, j = i+1, j+1 {
		tempStr[i], tempStr[j] = tempStr[j], tempStr[i]
	}
	output = string(tempStr)
	return output
}
