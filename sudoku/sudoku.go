package sudoku

import (
	"fmt"
)

type Sudoku struct {
	board     [9][9]int
	solutions [][9][9]int
}

func NewSudoku(board [9][9]int) *Sudoku {
	return &Sudoku{board: board}
}

func (s *Sudoku) possible(x int, y int, n int) bool {
	// Implementar
	return false
}

func (sudoku *Sudoku) Solve() bool {
	// Crea o reinicia las soluciones
	sudoku.solutions = [][9][9]int{}
	return sudoku.solve()
}

func (s *Sudoku) solve() bool {
	for y := range 9 {
		for x := range 9 {
			if s.board[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if s.possible(x, y, n) {
						s.board[y][x] = n
						s.solve() // if true, tengo al menos 1 solucion y puedo salir
						s.board[y][x] = 0
					}
				}
				return false
			}
		}
	}
	// Solucion correcta encontrada
	s.addSolution()
	return true
}

func (s *Sudoku) addSolution() {
	duplicate := [9][9]int{}
	for y := range 9 {
		for x := range 9 {
			duplicate[y][x] = s.board[y][x]
		}
	}

	s.solutions = append(s.solutions, duplicate)
}

func (s *Sudoku) PrintBoard() {
	printBoard(s.board)
}

func (s *Sudoku) PrintSolutions() {
	for _, solution := range s.solutions {
		printBoard(solution)
	}
}

func printBoard(board [9][9]int) {
	fmt.Println("┏━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━┓")
	for i, row := range board {
		fmt.Printf("┃ %v  %v  %v ┃ %v  %v  %v ┃ %v  %v  %v ┃\n",
			row[0], row[1], row[2],
			row[3], row[4], row[5],
			row[6], row[7], row[8])
		if i == 2 || i == 5 {
			fmt.Println("┣━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━┫")
		}
	}
	fmt.Println("┗━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━┛")
}
