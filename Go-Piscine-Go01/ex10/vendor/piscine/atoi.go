package piscine

func Atoi(s string) int {
	const intSize = 32 << (^uint(0) >> 63)

	if intSize == 32 && s == "-2147483648" {
		return -2147483648
	} else if intSize == 64 && s == "-9223372036854775808" {
		return -9223372036854775808
	}

	res := 0
	sign := 1
	for i, ch := range []byte(s) {
		if i == 0 && (ch == '+' || ch == '-') {
			if ch == '-' {
				sign = -1
			}
			continue
		}
		if ch < '0' || ch > '9' {
			return 0
		}
		res = res*10 + int(ch-'0')
	}
	return sign * res
}
