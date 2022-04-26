package piscine

func StrRev(s string) string {
	var res string
	for _, rune := range s {
		res = string(rune) + res
	}
	return res
}
