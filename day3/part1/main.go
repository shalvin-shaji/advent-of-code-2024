package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("Error trying to open file %x", err)
		fmt.Println("Cannot open the file")
		return
	}
    sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			number1, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("cannot convert number")
			}
			number2, err2 := strconv.Atoi(match[2])
			if err2 != nil {
				fmt.Println("cannot convert number")
			}
            sum += number1 * number2
		}
	}
    fmt.Println(sum)
}
