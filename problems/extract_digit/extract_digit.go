package main

import (
	"fmt"
	"strconv"
)

func extractDigitsString(n int) []int {

	s := strconv.Itoa(n)
	if n < 0 {
		s = s[1:]
	}

	digits := make([]int, len(s))

	for i, ch :=range s{
       digits[i] = int(ch - '0')
	}
   return digits
}

func main() {
	// Example usage
	numbers := []int{12345, -987, 0, 7}

	fmt.Println("\nUsing string conversion method (handles negatives):")
	for _, num := range numbers {
		digits := extractDigitsString(num)
		fmt.Printf("%d → %v\n", num, digits)
	}

}
