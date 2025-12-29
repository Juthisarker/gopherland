package isogram
import "strings"
func IsIsogram(word string) bool {
    seen := make(map[rune]bool)  // Create empty map
    
    upperWord := strings.ToUpper(word)  // "six-year-old" → "SIX-YEAR-OLD"
    
    for _, char := range upperWord {    // Loop: S, I, X, -, Y, E, A, R, -, O, L, D
        // Skip spaces and hyphens
        if char == ' ' || char == '-' {
            continue  
        }
        if seen[char] {
            return false
        }
        
        // Add to map: seen['S'] = true
        seen[char] = true
    }
    
    return true
}
