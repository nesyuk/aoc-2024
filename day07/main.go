package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2(result int, target int, counter *int, nums []int) {
	if result > target {
		return
	}
	if len(nums) == 0 {
		if result == target {
			*counter++
		}
		return
	}
	part2(result+nums[0], target, counter, nums[1:])
	part2(result*nums[0], target, counter, nums[1:])
	concat, err := strconv.Atoi(fmt.Sprintf("%d%d", result, nums[0]))
	if err != nil {
		panic(err)
	}
	part2(concat, target, counter, nums[1:])
}

func part1(result int, target int, counter *int, nums []int) {
	if result > target {
		return
	}
	if len(nums) == 0 {
		if result == target {
			*counter++
		}
		return
	}
	part1(result+nums[0], target, counter, nums[1:])
	part1(result*nums[0], target, counter, nums[1:])
}

func input(sc *bufio.Scanner) {
	targetsSum := 0
	for sc.Scan() {
		t := sc.Text()
		v := strings.Split(t, ": ")
		target, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		ops := make([]int, 0)
		opsT := strings.Split(v[1], " ")
		for _, opT := range opsT {
			op, err := strconv.Atoi(opT)
			if err != nil {
				panic(err)
			}
			ops = append(ops, op)
		}
		counter := 0
		part2(ops[0], target, &counter, ops[1:])
		if counter > 0 {
			targetsSum += target
		}
	}
	fmt.Println(targetsSum)
}

func main() {
	r, err := os.Open("./day07/data/input.txt")
	defer r.Close()
	if err != nil {
		panic(err)
	}
	input(bufio.NewScanner(r))
}
