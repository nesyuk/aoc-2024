package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pageSort(rules map[int]map[int]bool, update []int) func(i, j int) bool {
	return func(i, j int) bool {
		if rules[update[j]][update[i]] {
			return true
		}
		return false
	}
}

func part2(rules map[int]map[int]bool, updates [][]int) int {
	count := 0
	for _, update := range updates {
		if !sort.SliceIsSorted(update, pageSort(rules, update)) {
			sort.Slice(update, pageSort(rules, update))
			count += update[len(update)/2]
		}
	}
	return count
}

func part1(rules map[int]map[int]bool, updates [][]int) int {
	count := 0
	for _, update := range updates {
		if sort.SliceIsSorted(update, pageSort(rules, update)) {
			count += update[len(update)/2]
		}
	}
	return count
}

func input(sc *bufio.Scanner) (map[int]map[int]bool, [][]int) {
	rules := make(map[int]map[int]bool)
	updates := make([][]int, 0)
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			break
		}

		s := strings.Split(sc.Text(), "|")
		l, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		if _, exist := rules[r]; !exist {
			rules[r] = make(map[int]bool)
		}
		rules[r][l] = true
	}
	for sc.Scan() {
		s := strings.Split(sc.Text(), ",")
		pages := make([]int, 0)
		for _, v := range s {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			pages = append(pages, n)
		}
		updates = append(updates, pages)
	}
	return rules, updates
}

func main() {
	r, err := os.Open("./day05/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	rules, updates := input(sc)
	fmt.Println(part1(rules, updates))
	fmt.Println(part2(rules, updates))
}
