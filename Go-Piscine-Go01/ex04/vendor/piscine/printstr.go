package piscine

import "ft"

func PrintStr(s string) {
	for _, rune := range s {
		ft.PrintRune(rune)
	}
}
