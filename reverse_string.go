package reverse_string

func ReverseString(input string) (output string) {
	in := []rune(input)
	l := len(in)
	if l==0 {return}
	out := make([]rune, 0, l)
	for i := l-1; i >= 0; i--  {
		out = append(out, in[i])
	}
	output = string(out)
	return
}
