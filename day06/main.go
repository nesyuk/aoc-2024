package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nextDir = map[rune]rune{'^': '>', '>': 'V', 'V': '<', '<': '^'}

func next(lab [][]rune, r, c int, dir rune, visited *map[string]rune) (int, int, rune) {
	nr, nc := r, c
	switch dir {
	case '^':
		nr--
	case '>':
		nc++
	case 'V':
		nr++
	case '<':
		nc--
	}
	if nr < 0 || nr == len(lab) || nc < 0 || nc == len(lab[0]) {
		return nr, nc, dir
	}
	if lab[nr][nc] == '#' || lab[nr][nc] == 'O' {
		nr, nc = r, c
		dir = nextDir[dir]
		switch dir {
		case '^':
			nr--
		case '>':
			nc++
		case 'V':
			nr++
		case '<':
			nc--
		}
		if nr < 0 || nr == len(lab) || nc < 0 || nc == len(lab[0]) {
			return nr, nc, dir
		}
	}
	r, c = nr, nc
	s := fmt.Sprintf("%d-%d", r, c)
	(*visited)[s] = '+'
	return r, c, dir
}

func part2(lab [][]rune, row, col int, dir rune) int {
	loops := 0
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[0]); j++ {
			if lab[i][j] == '.' {
				lab[i][j] = 'O'

				visited := map[string]rune{}
				visited[fmt.Sprintf("%d-%d", row, col)] = dir
				r, c, d := row, col, dir
				for r >= 0 && c >= 0 && r < len(lab) && c < len(lab[0]) {
					if lab[r][c] == '#' || lab[r][c] == 'O' {
						// step back
						switch d {
						case '^':
							r++
						case '>':
							c--
						case 'V':
							r--
						case '<':
							c++
						}
						// turn
						d = nextDir[d]
					}
					switch d {
					case '^':
						r--
					case '>':
						c++
					case 'V':
						r++
					case '<':
						c--
					}

					s := fmt.Sprintf("%d-%d", r, c)
					if visitedDir, exists := visited[s]; exists {
						if visitedDir == d {
							loops++
							lab[i][j] = '.'
							break
						}
					}
					visited[s] = d

				}

				lab[i][j] = '.'
			}
		}
	}
	return loops
}

func part1(lab [][]rune, r, c int, dir rune) int {
	visited := map[string]rune{}
	visited[fmt.Sprintf("%d-%d", r, c)] = dir
	for r >= 0 && c >= 0 && r < len(lab) && c < len(lab[0]) {
		if lab[r][c] == '#' {
			// step back
			switch dir {
			case '^':
				r++
			case '>':
				c--
			case 'V':
				r--
			case '<':
				c++
			}
			s := fmt.Sprintf("%d-%d", r, c)
			visited[s] = '+'
			// turn
			dir = nextDir[dir]
		} else {
			s := fmt.Sprintf("%d-%d", r, c)
			if dir == '^' || dir == 'V' {
				visited[s] = '|'
			} else {
				visited[s] = '-'
			}
		}
		switch dir {
		case '^':
			r--
		case '>':
			c++
		case 'V':
			r++
		case '<':
			c--
		}
	}
	return len(visited)
}

func input(sc *bufio.Scanner) ([][]rune, int, int, rune) {
	lab := make([][]rune, 0)
	var row, col int
	var dir rune
	r := 0
	for sc.Scan() {
		t := sc.Text()
		for _, s := range []string{"^", ">", "V", "<"} {
			if strings.Contains(t, s) {
				row = r
				col = strings.Index(t, s)
				dir = []rune(s)[0]
			}
		}
		lab = append(lab, []rune(t))
		r++
	}
	return lab, row, col, dir
}

func main() {
	r, err := os.Open("./day06/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	lab, row, col, dir := input(sc)
	fmt.Println(part2(lab, row, col, dir))
}
