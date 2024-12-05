package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cut(a []int, i int) []int {
	return append(append([]int{}, a[:i]...), a[i+1:]...)
}

func isSafeCorrected(level []int) bool {
	firstDiff := level[1] - level[0]
	for i := 0; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]
		if diff == 0 || diff > 3 || diff < -3 || (firstDiff < 0 && diff > 0) || (firstDiff > 0 && diff < 0) {
			if i > 0 {
				return isSafe(cut(level, i-1)) || isSafe(cut(level, i)) || isSafe(cut(level, i+1))
			}
			return isSafe(cut(level, i)) || isSafe(cut(level, i+1))
		}
	}
	return true
}

func isSafe(level []int) bool {
	firstDiff := level[1] - level[0]
	for i := 0; i < len(level)-1; i++ {
		diff := level[i+1] - level[i]
		if diff == 0 || diff > 3 || diff < -3 || (firstDiff < 0 && diff > 0) || (firstDiff > 0 && diff < 0) {
			return false
		}
	}

	return true
}

func part1(sc *bufio.Scanner) int {
	safe := 0
	for sc.Scan() {
		arr := strings.Split(sc.Text(), " ")
		report := make([]int, len(arr))
		for i, s := range arr {
			n, _ := strconv.Atoi(s)
			report[i] = n
		}
		if isSafe(report) {
			safe++
		}
	}
	return safe
}

func part2(sc *bufio.Scanner) int {
	safe := 0
	for sc.Scan() {
		arr := strings.Split(sc.Text(), " ")
		report := make([]int, len(arr))
		for i, s := range arr {
			n, _ := strconv.Atoi(s)
			report[i] = n
		}
		if isSafeCorrected(report) {
			safe++
		}
	}
	return safe
}

func main() {
	r, _ := os.Open("./day02/data/input.txt")
	defer r.Close()
	sc := bufio.NewScanner(r)
	_ = sc
	fmt.Println(part1(sc))
	// fmt.Println(part2(sc))
}
