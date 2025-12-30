package luhn
import (
	"strings"
	"unicode"
)
func Valid(id string) bool {
    clean := strings.ReplaceAll(id, " ", "")
   	if len(clean) <= 1 {
		return false
	}
    for _, ch := range clean {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
    sum := 0
    shouldDouble := false
    	for i := len(clean) - 1; i >= 0; i-- {
		digit := int(clean[i] - '0') 
		
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		
		sum += digit
		shouldDouble = !shouldDouble  
	}

	return sum%10 == 0
}
