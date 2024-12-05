package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func distance(n, m int) int {
	if n > m {
		return n - m
	}
	return m - n
}

func part2(l, r []int) int {
	lmap, rmap := make(map[int]int), make(map[int]int)
	for i := range l {
		lmap[l[i]]++
		rmap[r[i]]++
	}
	totalDistance := 0
	for n := range lmap {
		totalDistance += lmap[n] * (n * rmap[n])
	}
	return totalDistance
}

func part1(ns, ms []int) int {
	slices.Sort(ns)
	slices.Sort(ms)
	totalDistance := 0
	for i, n := range ns {
		totalDistance += distance(n, ms[i])
	}
	return totalDistance
}

func input(f string) ([]int, []int) {
	r, _ := os.Open(f)
	defer r.Close()
	sc := bufio.NewScanner(r)

	ns, ms := make([]int, 0), make([]int, 0)
	var n, m int
	for sc.Scan() {
		if _, err := fmt.Sscanf(sc.Text(), "%d %d", &n, &m); err != nil {
			panic(err)
		}
		ns = append(ns, n)
		ms = append(ms, m)
	}
	return ns, ms
}

func main() {
	fmt.Println(part1(input("./day01/data/input.txt")))
	fmt.Println(part2(input("./day01/data/input.txt")))
}
