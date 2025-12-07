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
	} else if min == max && !isValidID(strconv.Itoa(min)) {
		return min
	} else {
		acc := 0
		for id := min; id<=max; id++ {
			if !isValidID(strconv.Itoa(id)) {
				acc += id
			}
		}
		return acc
	}
}

func isValidID(ID string) bool {
	return !bruteForceHasPattern(ID)
}

func bruteForceHasPattern(id string) bool {
	n := len(id)
	for size:=1; size <= n/2; size++ {
		if n % size != 0 {
			continue
		}

		pattern := id[:size]
		valid := true

		for i := size; i < n; i += size {
			if id[i:i+size] != pattern {
				valid = false
				break
			}
		}
		
		if valid {
			return true
		}
	}
	return false
}