package FamousProblemsWithGo

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	possibleBoards := SolveNQueens(4)

	if possibleBoards[0][0] != ".Q.." && possibleBoards[0][0] != "..Q." {
		t.Errorf("4x4 Queen Board solution seems to be wrong")
		PrintBoards(&possibleBoards)
	}

	possibleBoards = SolveNQueens(1)
	if possibleBoards[0][0] != "Q" {
		t.Errorf("1x1 Queen Board solution seems to be wrong")
		PrintBoards(&possibleBoards)
	}
}

func PrintBoards(boards *[][]string) {
	fmt.Printf("\n Print Board \n")
	for i := 0; i < len(*boards); i++ {
		for j := 0; j < len((*boards)[i]); j++ {
			fmt.Printf("%s ", (*boards)[i][j])
		}
		fmt.Printf("\n")
	}
}