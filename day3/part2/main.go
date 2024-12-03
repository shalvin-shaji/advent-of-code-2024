package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`(do\(\))|(don't\(\))|(mul\((\d+),(\d+)\))`)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("Error trying to open file %x", err)
		fmt.Println("Cannot open the file")
		return
	}
	sum := 0
	scanner := bufio.NewScanner(file)
	enable := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			main_match := match[0]
			if strings.HasPrefix(main_match, "don't") {
				enable = false
			} else if strings.HasPrefix(main_match, "do") {
				enable = true
			} else if strings.HasPrefix(main_match, "mul") && enable {
				number1, err := strconv.Atoi(match[4])
				if err != nil {
					fmt.Println("Error converting number")
				}
				number2, err := strconv.Atoi(match[5])
				if err != nil {
					fmt.Println("Error converting number")
				}
				sum += number1 * number2
			}

		}
	}
	fmt.Println(sum)
}
