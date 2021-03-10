// Additional math funcions
package math

import "time"

// Calculate a Fibonacci number
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	time.Sleep(0)

	return Fibonacci(n-1) + Fibonacci(n-2)
}
