// https://leetcode.com/problems/power-of-two

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}