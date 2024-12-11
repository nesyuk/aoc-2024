package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func pow10(n int) int {
	p := 1
	for i := 0; i < n; i++ {
		p *= 10
	}
	return p
}

func turn(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(fmt.Sprintf("%d", stone))%2 == 0 {
		p := pow10(len(fmt.Sprintf("%d", stone)) / 2)
		return []int{stone / p, stone % p}
	} else {
		return []int{stone * 2024}
	}
}

// time is a current time, starting from 0, times is a total number of iterations
func blink(stone int, time int, times int, memo *map[int]map[int]int) int {
	if time == times {
		return 1
	}
	if iters, exist := (*memo)[stone]; exist {
		if l, exist := iters[times-time+1]; exist {
			return l
		}
	}
	n := 0
	for _, s := range turn(stone) {
		d := blink(s, time+1, times, memo)
		if math.MaxInt-n < d {
			panic(d)
		}
		n += d
	}
	if _, exist := (*memo)[stone]; !exist {
		(*memo)[stone] = make(map[int]int)
	}
	(*memo)[stone][times-time+1] = n
	return n
}

func part2(stones []int, times int) int {
	sum := 0
	memo := make(map[int]map[int]int)
	for _, stone := range stones {
		sum += blink(stone, 0, times, &memo)
	}
	return sum
}

func part1(stones []int, times int) int {
	if times == 0 {
		return len(stones)
	}
	next := make([]int, 0)
	for _, s := range stones {
		if s == 0 {
			next = append(next, 1)
		} else if len(fmt.Sprintf("%d", s))%2 == 0 {
			p := pow10(len(fmt.Sprintf("%d", s)) / 2)
			next = append(next, s/p)
			next = append(next, s%p)
		} else {
			next = append(next, s*2024)
		}
	}
	return part1(next, times-1)
}

func input(sc *bufio.Scanner) []int {
	stones := make([]int, 0)
	sc.Scan()
	vals := strings.Split(sc.Text(), " ")
	for _, v := range vals {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		stones = append(stones, n)
	}
	return stones
}

func main() {
	r, err := os.Open("./day11/data/input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(r)
	//fmt.Println(part1(input(sc), 25))
	fmt.Println(part2(input(sc), 75))
}
