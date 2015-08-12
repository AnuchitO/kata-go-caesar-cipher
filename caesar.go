package main

func Caesar(text string, shift int) string {
	b := []byte(text)

	if len(b) == 1 {
		return string(b[0] + byte(shift))
	}

	if len(b) == 2 {
		b[0] = b[0] + byte(shift)
		b[1] = b[1] + byte(shift)

		return string(b)
	}

	b[0] = b[0] + byte(shift)
	b[1] = b[1] + byte(shift)
	b[2] = b[2] + byte(shift)

	return string(b)
}
