package reverse_string

func ReverseString(input string) (output string) {
	rn := []rune(input)
	for i:=len(rn)-1; i>=0; i--{
		output += string(rn[i])
	}
	return output
}

