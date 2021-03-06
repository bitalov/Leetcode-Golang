// https://leetcode.com/problems/optimal-division

func optimalDivision(nums []int) string {
   strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}

	switch len(strs) {
	case 1:
		return strs[0]
	case 2:
		return strings.Join(strs, "/")
	default:
		return strs[0] + "/(" + strings.Join(strs[1:], "/") + ")"
	} 
}