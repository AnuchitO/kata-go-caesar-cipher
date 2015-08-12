package main

func Caesar(text string, shift int) string {
	b := []byte(text)
	return string(b[0] + byte(shift))
}
