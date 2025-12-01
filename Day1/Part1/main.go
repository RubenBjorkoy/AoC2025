package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
);
const fileLength int = 4424

func main() {
	file, err := os.Open("input")
	defer file.Close()

	r := bufio.NewReader(file)

	if err != nil {
		panic(err)
	}

	i := 0
	acc := 0
	currValue := 50

	for i < fileLength {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			currValue = rotateDial(string(line), currValue)
			if currValue == 0 {acc++}
			i++
		}
		if err != nil {
			break;
		}
	}
	fmt.Println(acc)
}


func rotateDial(instruction string, current int) int {
	ins, _ := strconv.Atoi(instruction[1:])
	if string(instruction[0]) ==  "L" {
		current -= ins
		for current < 0 {
			current += 100
		}
	} else {
		current += ins
				for current > 99 {
			current -= 100
		}
	}

	return current
}


// func readFile(FileName string) [fileLength]string {
// 	file, err := os.Open(FileName)
// 	defer file.Close()

// 	r := bufio.NewReader(file)

// 	if err != nil {
// 		panic(err)
// 	}

// 	instructions := [fileLength]string{}
// 	i := 0

// 	for i < fileLength {
// 		line, _, err := r.ReadLine()
// 		if len(line) > 0 {
// 			instructions[i] = string(line)
// 			i++
// 		}
// 		if err != nil {
// 			break;
// 		}
// 	}
// 	return instructions
// }