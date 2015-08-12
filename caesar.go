package main

func Caesar(text string, shift int32) string {
	var b []byte

	for _, c := range text {
		b = append(b, byte('A'+(c+shift-'A')%26))
	}

	return string(b)
}
