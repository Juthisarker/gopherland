package microblog


func Truncate(phrase string) string {
	runes := []rune(phrase) // convert to Unicode code points

	if len(runes) <= 5 {
		return phrase
	}

	return string(runes[:5])
}
