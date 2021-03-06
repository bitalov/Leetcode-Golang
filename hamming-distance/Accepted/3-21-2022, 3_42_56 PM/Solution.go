// https://leetcode.com/problems/hamming-distance

func hammingDistance(x int, y int) int {
	x ^= y

	res := 0
	for x > 0 {
		res += x & 1
		x >>= 1
	}

	return res
}