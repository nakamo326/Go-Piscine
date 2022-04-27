package piscine

func isInRange(res, sign int, ch byte) bool {
	const intSize = 32 << (^uint(0) >> 63)
	var div int
	if intSize == 32 {
		div = 214748364
	} else {
		div = 922337203685477580
	}
	if (div < res || (div == res && ch > '7')) && sign > 0 {
		return false
	}
	if (div < res || (div == res && ch > '8')) && sign < 0 {
		return false
	}
	return true
}

func Atoi(s string) int {
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
		if !isInRange(res, sign, ch) {
			return 0
		}
		res = res*10 + int(ch-'0')
	}
	return sign * res
}
