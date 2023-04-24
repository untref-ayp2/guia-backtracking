package sudoku

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func (s *Sudoku) VisualSolve(p *widgets.Paragraph, delay time.Duration, steps *int) bool {
	for y := range 9 {
		for x := range 9 {
			if s.board[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if s.possible(x, y, n) {
						*steps++
						s.board[y][x] = n
						s.Render(p, x, y, *steps)
						time.Sleep(delay)
						if s.VisualSolve(p, delay, steps) {
							s.Render(p, -1, -1, *steps)
							return true
						}
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

func (s *Sudoku) Render(p *widgets.Paragraph, currentX int, currentY int, steps int) {
	print := "┏━━━━━━━┳━━━━━━━┳━━━━━━━┓\n"
	for row := range 9 {
		print += "┃ "
		for col := range 9 {
			if col == 3 || col == 6 {
				print += "┃ "
			}
			if s.board[row][col] == 0 {
				print += "[0](fg:red) "
			} else if row == currentY && col == currentX {
				print += fmt.Sprintf("[%d](fg:blue) ", s.board[row][col])
			} else {
				print += fmt.Sprintf("%d ", s.board[row][col])
			}

			if col == 8 {
				print += "┃"
			}
		}
		if row == 2 || row == 5 {
			print += "\n┣━━━━━━━╋━━━━━━━╋━━━━━━━┫\n"
		} else {
			print += "\n"
		}
	}
	print += "┗━━━━━━━┻━━━━━━━┻━━━━━━━┛\n"

	print += fmt.Sprintf("\n\nPasos: [%d](fg:yellow) ", steps)

	if currentX == -1 || currentY == -1 {
		print += "\n\nPresionar ['q'](fg:blue) para finalizar."
	}

	p.Text = print
	ui.Render(p)
}
