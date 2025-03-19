
// This function takes an integer n as input and returns its factorial.
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}