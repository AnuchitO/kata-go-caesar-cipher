package main

func Caesar(text string, shift int32) string {
	var b []byte

	for _, c := range text {
		b = append(b, byte('A'+(c+shift-'A')%26))
	}

	return string(b)
}

type CaesarCipher struct {
	text  string
	shift int32
}

func (caesar *CaesarCipher) GetCipher() string {
	return Caesar(caesar.text, caesar.shift)
}
