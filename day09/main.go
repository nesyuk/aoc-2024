package main

import (
	"bufio"
	"fmt"
	"os"
)

func decode(encoded []int) []int {
	decoded := make([]int, 0)
	fileId := 0
	for i, length := range encoded {
		if i%2 == 0 {
			// file
			for j := 0; j < length; j++ {
				decoded = append(decoded, fileId)
			}
			fileId++
		} else {
			// space
			for j := 0; j < length; j++ {
				decoded = append(decoded, -1)
			}
		}
	}
	return decoded
}

func checkSum(decoded []int) uint {
	sum := uint(0)
	for i := 0; i < len(decoded); i++ {
		if decoded[i] != -1 {
			sum += uint(i) * uint(decoded[i])
		}
	}
	return sum
}

func part1(encoded []int) uint {
	decoded := decode(encoded)
	for i, j := 0, len(decoded)-1; i < j; {
		if decoded[i] != -1 {
			i++
		} else if decoded[j] == -1 {
			j--
		} else {
			// swap
			decoded[i], decoded[j] = decoded[j], decoded[i]
		}
	}
	return checkSum(decoded)
}

func swap(decoded []int, to, from int, length int) []int {
	for l := 0; l < length; l++ {
		decoded[to] = decoded[from]
		decoded[from] = -1
		to++
		from--
	}
	return decoded
}

func part2(encoded []int) uint {
	decoded := decode(encoded)
	last := len(decoded) - 1
	for last > 0 {
		j := last
		for ; j >= 0 && decoded[j] == -1; j-- {
		}
		if j == 0 {
			return checkSum(decoded)
		}
		jLen := 0
		for i := j - 1; i >= 0 && decoded[i] == decoded[j]; i-- {
			jLen++
		}
		for i := 0; i < j-jLen; {
			for ; i < j-jLen && decoded[i] != -1; i++ {
			}
			iLen := 0
			for k := i + 1; k < j-jLen && decoded[k] == decoded[i]; k++ {
				iLen++
			}
			if i+iLen < j-jLen && iLen >= jLen {
				decoded = swap(decoded, i, j, jLen+1)
				last = j - jLen - 1
				break
			} else {
				i += iLen + 1
			}
		}
		last = j - jLen - 1
	}
	return checkSum(decoded)
}

func input(sc *bufio.Scanner) []int {
	sc.Scan()
	t := sc.Text()
	encoded := make([]int, 0)
	for _, r := range []rune(t) {
		encoded = append(encoded, int(r-'0'))
	}
	return encoded
}

func main() {
	r, err := os.Open("./day09/data/input.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Println(part1(input(bufio.NewScanner(r))))
	fmt.Println(part2(input(bufio.NewScanner(r))))
}
