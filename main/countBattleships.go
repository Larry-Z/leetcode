package src

import "fmt"

func countBattleships(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}
	count := 0
	for i, arr := range board {
		for j, cell := range arr {
			if cell == '.' {
				continue
			}
			if i > 0 && board[i-1][j] == 'X' {
				continue
			}
			if j > 0 && board[i][j-1] == 'X' {
				continue
			}
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countBattleships([][]byte{
		[]byte{'X', '.', '.', 'X'},
		[]byte{'.', '.', '.', 'X'},
		[]byte{'.', '.', '.', 'X'},
	}))
}
