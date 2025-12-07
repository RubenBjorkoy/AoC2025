package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
);
const fileLength int = 1

func main() {
	file, err := os.Open("input")
	defer file.Close()

	r := bufio.NewReader(file)

	if err != nil {
		panic(err)
	}

	invalidIdSum := 0
	i := 0

	for i < fileLength {
		line, _, err := r.ReadLine()
		if len(line) > 0 {

			ranges := strings.Split(string(line), ",")
			for _, idRange := range ranges {
				min, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
				max, _ := strconv.Atoi(strings.Split(idRange, "-")[1])
				invalidIdSum += loopRange(min, max)
			}

			i++
		}
		if err != nil {
			break;
		}
	}
	fmt.Println(invalidIdSum)
}

func loopRange(min int, max int) int {
	if max < min {
		return 0
	} else if min == max && !isValidID(min) {
		return min
	} else {
		acc := 0
		for id := min; id<=max; id++ {
			if !isValidID(id) {
				acc += id
			}
		}
		return acc
	}
	
	return 0
}

func isValidID(ID int) bool {
	stringID := strconv.Itoa(ID)
	if len(stringID) % 2 != 0 {
		return true
	}
	half1 := stringID[:(len(stringID) / 2)]
	half2 := stringID[(len(stringID) / 2):]
	return !(half1 == half2)
}