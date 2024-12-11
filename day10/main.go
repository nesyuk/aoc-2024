package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(plain [][]int, trailheads [][]int) int {
	scores := 0
	for _, start := range trailheads {
		score := 0
		visited := map[string]bool{}
		queue := [][]int{start}
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]
			if plain[r][c] == 9 {
				score++
				scores++
			} else {
				for _, adj := range [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
					r1, c1 := adj[0], adj[1]
					if r1 >= 0 && c1 >= 0 && r1 < len(plain) && c1 < len(plain[0]) {
						if !visited[fmt.Sprintf("%d-%d", r1, c1)] && plain[r1][c1]-plain[r][c] == 1 {
							queue = append(queue, []int{r1, c1})
							// uncomment for part1
							// visited[fmt.Sprintf("%d-%d", r1, c1)] = true
						}
					}
				}
			}
		}
	}
	return scores
}

func input(sc *bufio.Scanner) ([][]int, [][]int) {
	trailheads := make([][]int, 0)
	plain := make([][]int, 0)
	r := 0
	for sc.Scan() {
		plain = append(plain, make([]int, 0))
		for c, run := range []rune(sc.Text()) {
			height := int(run - '0')
			plain[r] = append(plain[r], height)
			if height == 0 {
				trailheads = append(trailheads, []int{r, c})
			}
		}
		r++
	}
	return plain, trailheads
}

func main() {
	r, err := os.Open("./day10/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	plain, trailheads := input(sc)
	fmt.Println(part1(plain, trailheads))
}
