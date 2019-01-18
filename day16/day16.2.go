package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	opcodes := []func([]int, int, int) int{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

	possibleCodes := make([][]bool, len(opcodes)*len(opcodes))
	for i := range possibleCodes {
		for j := range possibleCodes[i] {
			possibleCodes[i][j] = true
		}
	}
	fmt.Println(possibleCodes)
	before, op, after := readData("data1.txt")

	for i := range before {
		fmt.Println(i)
		for j := range opcodes {
			if !possibleCodes[j][op[i][0]] {
				continue
			}

			result := make([]int, len(before[i]))
			copy(result, before[i])
			result[op[i][3]] = opcodes[j](before[i], op[i][1], op[i][2])
			possibleCodes[j][op[i][0]] = comp(result, after[i])
		}
	}

	fmt.Println(possibleCodes)

}

func readData(fn string) ([][]int, [][]int, [][]int) {
	file, err := os.Open(fn)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	before := [][]int{}
	op := [][]int{}
	after := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		if scanner.Text()[:3] == "Bef" || scanner.Text()[:3] == "Aft" {
			regr := arrToi(strings.Split(scanner.Text()[9:len(scanner.Text())-1], ", "))
			if scanner.Text()[:3] == "Bef" {
				before = append(before, regr)
			} else {
				after = append(after, regr)
			}
			continue
		}
		op = append(op, arrToi(strings.Split(scanner.Text(), " ")))
	}

	return before, op, after
}

func arrToi(str []string) []int {
	conv := make([]int, len(str))

	for i := range str {
		conv[i], _ = strconv.Atoi(str[i])
	}

	return conv
}

func comp(rega []int, regb []int) bool {
	if len(rega) != len(regb) {
		return false
	}

	for i := range rega {
		if rega[i] != regb[i] {
			return false
		}
	}

	return true
}

func addr(regr []int, a int, b int) int {
	return regr[a] + regr[b]
}

func addi(regr []int, a int, b int) int {
	return regr[a] + b
}

func mulr(regr []int, a int, b int) int {
	return regr[a] * regr[b]
}

func muli(regr []int, a int, b int) int {
	return regr[a] * b
}

func banr(regr []int, a int, b int) int {
	return regr[a] & regr[b]
}

func bani(regr []int, a int, b int) int {
	return regr[a] & b
}

func borr(regr []int, a int, b int) int {
	return regr[a] | regr[b]
}

func bori(regr []int, a int, b int) int {
	return regr[a] | b
}

func setr(regr []int, a int, b int) int {
	return regr[a]
}

func seti(regr []int, a int, b int) int {
	return a
}

func gtir(regr []int, a int, b int) int {
	if a > regr[b] {
		return 1
	}
	return 0
}

func gtri(regr []int, a int, b int) int {
	if regr[a] > b {
		return 1
	}
	return 0
}

func gtrr(regr []int, a int, b int) int {
	if regr[a] > regr[b] {
		return 1
	}
	return 0
}

func eqir(regr []int, a int, b int) int {
	if a == regr[b] {
		return 1
	}
	return 0
}

func eqri(regr []int, a int, b int) int {
	if regr[a] == b {
		return 1
	}
	return 0
}

func eqrr(regr []int, a int, b int) int {
	if regr[a] == regr[b] {
		return 1
	}
	return 0
}
