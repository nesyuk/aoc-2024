package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(frequencies map[rune][][]int, w, h int) int {
	antinodes := map[string]bool{}
	for _, antennas := range frequencies {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				dr, dc := antennas[i][0]-antennas[j][0], antennas[i][1]-antennas[j][1]
				r, c := antennas[i][0]+dr, antennas[i][1]+dc
				if r >= 0 && r < h && c >= 0 && c < w {
					antinodes[fmt.Sprintf("%d-%d", r, c)] = true
				}
				r, c = antennas[j][0]-dr, antennas[j][1]-dc
				if r >= 0 && r < h && c >= 0 && c < w {
					antinodes[fmt.Sprintf("%d-%d", r, c)] = true
				}
			}
		}
	}
	return len(antinodes)
}

func part2(frequencies map[rune][][]int, w, h int) int {
	antinodes := map[string]bool{}
	for _, antennas := range frequencies {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				dr, dc := antennas[i][0]-antennas[j][0], antennas[i][1]-antennas[j][1]
				for r, c := antennas[i][0], antennas[i][1]; r >= 0 && c >= 0 && c < w; r, c = r+dr, c+dc {
					antinodes[fmt.Sprintf("%d-%d", r, c)] = true
				}
				for r, c := antennas[i][0], antennas[i][1]; r < h && c >= 0 && c < w; r, c = r-dr, c-dc {
					antinodes[fmt.Sprintf("%d-%d", r, c)] = true
				}
			}
		}
	}
	return len(antinodes)
}

func input(sc *bufio.Scanner) (map[rune][][]int, int, int) {
	frequencies := map[rune][][]int{}
	w, h := 0, 0
	for sc.Scan() {
		h++
		t := sc.Text()
		if w == 0 {
			w = len([]rune(t))
		}
		for c, r := range t {
			if r != '.' {
				if _, exist := frequencies[r]; !exist {
					frequencies[r] = [][]int{}
				}
				frequencies[r] = append(frequencies[r], []int{h - 1, c})
			}
		}
	}
	return frequencies, w, h
}

func main() {
	r, err := os.Open("./day08/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	frequencies, w, h := input(sc)
	// fmt.Println(part1(frequencies, w, h))
	fmt.Println(part2(frequencies, w, h))
}
