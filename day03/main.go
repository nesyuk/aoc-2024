package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var op2 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(don't\(\))|(do\(\))`)
var op1 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func part2(sc *bufio.Scanner) {
	sum := 0
	on := true
	for sc.Scan() {
		str := sc.Text()
		//fmt.Println(str)
		matches := op2.FindAllStringSubmatch(str, -1)
		for _, m := range matches {
			if m[0] == "do()" {
				fmt.Println("do()")
				on = true
			} else if m[0] == "don't()" {
				fmt.Println("don't()")
				on = false
			} else if on {
				fmt.Println(m)
				a, b := m[1], m[2]
				n1, err := strconv.Atoi(a)
				if err != nil {
					panic(a)
				}
				n2, err := strconv.Atoi(b)
				if err != nil {
					panic(b)
				}
				sum += n1 * n2
			}
		}
	}
	fmt.Println(sum)
}

func scan(str string) int {
	sum := 0
	matches := op1.FindAllStringSubmatch(str, -1)
	for _, m := range matches {
		if m[0] == "do()" {
			fmt.Println("do()")
		} else if m[0] == "don't()" {
			fmt.Println("don't()")
		} else {
			fmt.Println(m)
			a, b := m[1], m[2]
			n1, err := strconv.Atoi(a)
			if err != nil {
				panic(a)
			}
			n2, err := strconv.Atoi(b)
			if err != nil {
				panic(b)
			}
			sum += n1 * n2
		}
	}
	return sum
}

func part1(sc *bufio.Scanner) {
	sum := 0
	for sc.Scan() {
		str := sc.Text()
		matches := op1.FindAllStringSubmatch(str, -1)
		for _, m := range matches {
			a, b := m[1], m[2]
			n1, err := strconv.Atoi(a)
			if err != nil {
				panic(a)
			}
			n2, err := strconv.Atoi(b)
			if err != nil {
				panic(b)
			}
			sum += n1 * n2
		}
	}
	fmt.Println(sum)
}

func main() {
	f, err := os.Open("./day03/data/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	part1(sc) // 180233229
	//part2(sc) // 180233229
}
