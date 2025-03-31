package iteration

import "strings"

func Repeated (charac string, repeat int)string {
	var repeated strings.Builder
	for i:= 0 ; i< repeat ; i++ {
		repeated.WriteString(charac)
	}
	return repeated.String()
}