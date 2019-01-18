package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	changesStr := strings.Split(strings.Trim(string(data), "\n "), "\n")
	changes := make([]int, len(changesStr))

	for i := range changesStr {
		changes[i], err = strconv.Atoi(changesStr[i])
		if err != nil {
			panic(err)
		}
	}

	frequency := 0

	for i := range changes {
		frequency += changes[i]
	}

	fmt.Println(frequency)

	frequency = 0
	frequencyLog := []int{frequency}
	i := 0

	for {
		frequency += changes[i%len(changes)]
		if intInSlice(frequency, frequencyLog) {
			break
		}
		frequencyLog = append(frequencyLog, frequency)
		i++
	}

	fmt.Println(i, frequency)
}

func intInSlice(intToFind int, list []int) bool {
	for i := range list {
		if list[i] == intToFind {
			return true
		}
	}
	return false
}
