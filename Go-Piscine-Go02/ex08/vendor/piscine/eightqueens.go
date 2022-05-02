package piscine

import "ft"

func printAnswer(answer [8]int) {
	for _, v := range answer {
		ft.PrintRune(rune('1' + v))
	}
	ft.PrintRune('\n')
}

func setNumDir(board *[8][8]int, i, j, num int, dir []int) {
	x := i
	y := j
	for ; (0 <= x && x < 8) && (0 <= y && y < 8); x, y = x+dir[0], y+dir[1] {
		board[x][y] += num
	}
}

func setBoard(board *[8][8]int, i, j, num int) {
	var k int
	for k = 0; k < 8; k++ {
		board[i][k] += num
		board[k][j] += num
	}
	dirList := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, dir := range dirList {
		setNumDir(board, i, j, num, dir)
	}
}

func solve(answer [8]int, board [8][8]int, i int) {
	if i == 8 {
		printAnswer(answer)
		return
	}

	for j := 0; j < 8; j++ {
		if board[i][j] == 0 {
			answer[i] = j
			setBoard(&board, i, j, 1)
			solve(answer, board, i+1)
			setBoard(&board, i, j, -1)
		}
	}
}

func EightQueens() {
	var answer [8]int
	var board [8][8]int

	solve(answer, board, 0)
}
