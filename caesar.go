package main

func Caesar(text string, shift int) string {
	b := []byte(text)

	for index := 0; index < len(b); index++ {
		b[index] = b[index] + byte(shift)
	}

	return string(b)
}
