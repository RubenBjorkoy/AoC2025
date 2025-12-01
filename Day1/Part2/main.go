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
			currValue, acc = rotateDial(string(line), currValue, acc)
			i++
		}
		if err != nil {
			break;
		}
	}
	fmt.Println(acc)
}


func rotateDial(instruction string, current int, acc int) (int, int) {
    amount, _ := strconv.Atoi(instruction[1:])
    dir := instruction[0]

    if dir == 'L' {
        for i := 0; i < amount; i++ {
            current--
            if current < 0 {
                current = 99
            }
            if current == 0 {
                acc++
            }
        }
    } else {
        for i := 0; i < amount; i++ {
            current++
            if current > 99 {
                current = 0
            }
            if current == 0 {
                acc++
            }
        }
    }

    return current, acc
}