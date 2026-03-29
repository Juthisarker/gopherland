package atbashcipher
import "strings"
func Atbash(s string) string {
		var result []rune

	for _, ch := range strings.ToLower(s) {
		if ch >= 'a' && ch <= 'z' {
			result = append(result, 'a'+('z'-ch))
		} else if ch >= '0' && ch <= '9' {
			result = append(result, ch)
		}
	}

	// group into chunks of 5
	var groups []string
	for i := 0; i < len(result); i += 5 {
		end := i + 5
		if end > len(result) {
			end = len(result)
		}
		groups = append(groups, string(result[i:end]))
	}

	return strings.Join(groups, " ")
}
