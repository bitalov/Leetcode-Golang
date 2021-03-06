// https://leetcode.com/problems/add-strings

func addStrings(s1 string, s2 string) string {
    if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	n1, n2 := len(s1), len(s2)
	a1, a2 := []byte(s1), []byte(s2)

	carry := byte(0)

	buf := make([]byte, n2+1)
	buf[0] = '1'

	i := 1
	for i <= n2 {
		if i <= n1 {
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry

		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])
}