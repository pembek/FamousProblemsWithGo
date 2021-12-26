package FamousProblemsWithGo

//The n-queens puzzle is the problem of placing n queens on an n x n
//chessboard such that no two queens attack each other.
//
//Given an integer n, return all distinct solutions to the n-queens puzzle.
//You may return the answer in any order.
//
//Each solution contains a distinct board configuration of the n-queens'
//placement, where 'Q' and '.' both indicate a queen and an empty space,
//respectively.

// SolveNQueens Input: n = 4
//Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
func SolveNQueens(n int) [][]string {
	board := []string{}

	for v := 0; v < n; v++ {
		board = append(board,"")
		for k := 0; k < n; k++ {
			board[v] = board[v] + "."
		}
	}

	var res [][]string
	Solve(&res, &board, 0)

	//Start in the leftmost columm
	//If all queens are placed, return true
	//for (every possible choice among the rows in this column)
	//if the queen can be placed safely there,
	//make that choice and then recursively try to place the rest of the queens if recursion successful, return true
	//if !successful, remove queen and try another row in this column
	//if all rows have been tried and nothing worked, return false to trigger backtracking

	return res
}

func Solve(res *[][]string, board *[]string, row int) {
	//fmt.Printf("\nSolve, ROW %d\n", row)
	if row == len(*board) {
		*res = append(*res, construct(board))
		return
	}

	for col := 0; col < len(*board); col++ {
		if IsValid(board, row, col) {

			//for j := 0; j < len(*board); j++ {
			//	fmt.Printf("%s , ", (*board)[j])
			//}

			(*board)[row] = (*board)[row][:col] + "Q" + (*board)[row][col+1:]
			Solve(res, board, row + 1)
			(*board)[row] = (*board)[row][:col] + "." + (*board)[row][col+1:]
		}
	}
}

func construct(board *[]string) []string {
	var path []string

	for i:= 0; i < len(*board); i++ {
		path = append(path, (*board)[i])
	}
	return path
}

func IsValid(board *[]string, row int, col int) bool {
	//fmt.Printf("\nIsValid, ROW %d , COL %d, 00: %c\n", row, col, (*board)[0][0])
	for i := 0; i < row; i++ { // check Queens in same column
		if (*board)[i][col] == 'Q' {
			//fmt.Printf("Queen on the same column")
			return false
		}
	}

	//only need to check backwards, because of backtracking
	for i, j := row, col; i >= 0 && j >= 0; {
		if (*board)[i][j] == 'Q' {
			//fmt.Printf("Queen on left top diagonal")
			return false
		}
		i--
		j--
	}

	//only need to check backwards, because of backtracking
	for i, j := row, col; j < len(*board) && i >= 0; {
		if (*board)[i][j] == 'Q' {
			//fmt.Printf("Queen on right top diagonal")
			return false
		}
		i--
		j++
	}
	return true
}
