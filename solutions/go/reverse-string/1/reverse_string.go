package reverse

func Reverse(input string) string {
     // Convert to a slice of runes to properly handle Unicode.
    runes:= []rune(input)
    //reverse the slice in place
    for i,j:= 0, len(runes) - 1; i < j; i,j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
