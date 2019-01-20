package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inst struct {
	opcode func([]int, int, int) int
	a      int
	b      int
	c      int
}

func main() {
	opcodes := []func([]int, int, int) int{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr, modr}
	opcodesStr := []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr", "modr"}

	ip, ops := parseFile("data.txt", opcodes, opcodesStr)

	for initValue := 0; initValue <= 1; initValue++ {
		regr := make([]int, 6)
		regr[0] = initValue
		ipp := &regr[ip]
		for *ipp < len(ops) {
			regr[ops[*ipp].c] = ops[*ipp].opcode(regr, ops[*ipp].a, ops[*ipp].b)
			*ipp++
		}
		*ipp--
		fmt.Printf("Part %v: %v\n", initValue+1, regr)
	}
}

func parseFile(fileName string, opCodes []func([]int, int, int) int, opCodesStr []string) (int, []inst) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	ip := atoi(strings.Split(scanner.Text(), " ")[1])

	ops := []inst{}

	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		f := opCodes[sliceFind(opCodesStr, op[0])]
		ops = append(ops, inst{f, atoi(op[1]), atoi(op[2]), atoi(op[3])})
	}
	return ip, ops
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func sliceFind(arr []string, toFind string) int {
	for i := range arr {
		if arr[i] == toFind {
			return i
		}
	}
	return -1
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

func modr(regr []int, a int, b int) int {
	return regr[a] % regr[b]
}
