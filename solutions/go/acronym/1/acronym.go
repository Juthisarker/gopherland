// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym
import (
	"strings"
	"unicode"
)
// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

    s= strings.ReplaceAll(s,"-"," ")
    var acronym []rune
    newWord:=true

    for _, r:= range s {
        if unicode.IsLetter(r){
            if newWord {
                acronym = append(acronym, unicode.ToUpper(r))
                newWord = false
                }
	} else if r == '\'' {
		continue
	} else {
		newWord = true
	}
        }
        
	return string(acronym)
}
