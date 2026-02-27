package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}	
  reversed := 0
  for x > reversed {
	reversed = reversed * 10 + x%10
	x /= 10
  }

  return x == reversed || x == reversed/10
}

func main() {
	// Example usage
	numbers := []int{121, -121, 10, 12321, 0}
	for _, num := range numbers {
		fmt.Printf("isPalindrome(%d) = %t\n", num, isPalindrome(num))
	}
}