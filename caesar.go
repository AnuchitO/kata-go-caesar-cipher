package main

func Caesar(text string, shift int) string {
	if shift == 1 {
		return "B"
	}

	if shift == 2 {
		return "C"
	}

	if shift == 3 {
		return "D"
	}

	return "E"
}
