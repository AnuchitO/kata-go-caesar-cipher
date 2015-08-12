package main

func Caesar(text string, shift int) string {
	b := []byte(text)

	for index := 0; index < len(b); index++ {
		b[index] = byte(65) + (b[index]+byte(shift)-65)%26
	}

	return string(b)
}
