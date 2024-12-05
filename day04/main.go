package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readMatrix(sc *bufio.Scanner) [][]string {
	m := make([][]string, 0)
	for sc.Scan() {
		m = append(m, strings.Split(sc.Text(), ""))
	}
	return m
}

func part1(m [][]string) int {
	const xmas = "XMAS"
	counter := 0
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if m[r][c] == "X" {
				// up
				i, j := r, c
				for k, char := range xmas {
					if i < 0 || m[i][j] != string(char) {
						break
					}
					i--
					if k == len(xmas)-1 {
						counter++
					}
				}
				// down
				i, j = r, c
				for k, char := range xmas {
					if i > len(m)-1 || m[i][j] != string(char) {
						break
					}
					i++
					if k == len(xmas)-1 {
						counter++
					}
				}
				// right
				i, j = r, c
				for k, char := range xmas {
					if j > len(m[0])-1 || m[i][j] != string(char) {
						break
					}
					j++
					if k == len(xmas)-1 {
						counter++
					}
				}
				// left
				i, j = r, c
				for k, char := range xmas {
					if j < 0 || m[i][j] != string(char) {
						break
					}
					j--
					if k == len(xmas)-1 {
						counter++
					}
				}
				// up-right
				i, j = r, c
				for k, char := range xmas {
					if i < 0 || j > len(m[0])-1 || m[i][j] != string(char) {
						break
					}
					j++
					i--
					if k == len(xmas)-1 {
						counter++
					}
				}
				// up-left
				i, j = r, c
				for k, char := range xmas {
					if i < 0 || j < 0 || m[i][j] != string(char) {
						break
					}
					j--
					i--
					if k == len(xmas)-1 {
						counter++
					}
				}
				// down-right
				i, j = r, c
				for k, char := range xmas {
					if i > len(m)-1 || j > len(m[0])-1 || m[i][j] != string(char) {
						break
					}
					j++
					i++
					if k == len(xmas)-1 {
						counter++
					}
				}
				// down-left
				i, j = r, c
				for k, char := range xmas {
					if i > len(m)-1 || j < 0 || m[i][j] != string(char) {
						break
					}
					j--
					i++
					if k == len(xmas)-1 {
						counter++
					}
				}
			}
		}
	}
	return counter
}

func part2(m [][]string) int {
	counter := 0
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if m[r][c] == "A" {
				if c > 0 && c < len(m)-1 && r > 0 && r < len(m[0])-1 {
					upRight := m[r-1][c+1]
					upLeft := m[r-1][c-1]
					downRight := m[r+1][c+1]
					downLeft := m[r+1][c-1]

					if (upRight == "M" && downLeft == "S" || upRight == "S" && downLeft == "M") &&
						(upLeft == "M" && downRight == "S" || upLeft == "S" && downRight == "M") {
						counter++
					}
				}

			}
		}
	}
	return counter
}

func main() {
	r, err := os.Open("./day04/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	matrix := readMatrix(sc)
	fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}
