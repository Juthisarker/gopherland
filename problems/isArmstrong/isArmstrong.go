package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	if n < 0 {
		return false
	}

	temp := n
	digits := 0
	for temp > 0 {
		temp /= 10
		digits++
	}

	sum := 0
	temp = n

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == n
}

func main() {
	numbers := []int{153, 370, 371, 9474, 9475, -153, 0}
	for _, num := range numbers {
		fmt.Printf("isArmstrong(%d) = %t\n", num, isArmstrong(num))
	}
}
