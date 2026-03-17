package diffsquares

func SquareOfSum(n int) int {
    // Formula: (n*(n+1)/2)²
	sum := n * (n + 1) / 2
    return sum * sum
}

func SumOfSquares(n int) int {
	    // Formula: n*(n+1)*(2n+1)/6
    return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	  return SquareOfSum(n) - SumOfSquares(n)
}
