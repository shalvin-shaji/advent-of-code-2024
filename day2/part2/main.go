package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
func isDecreasing(intlist []int) bool {
	for i := 1; i < len(intlist); i++ {
		if intlist[i] > intlist[i-1] {
			return false
		}
	}
	return true
}
func isIncreasing(intlist []int) bool {
	for i := 1; i < len(intlist); i++ {
		if intlist[i] < intlist[i-1] {
			return false
		}
	}
	return true
}

func isSafe(intlist []int) bool {
	if len(intlist) < 1 {
		return false
	}
	if !isIncreasing(intlist) && !isDecreasing(intlist) {
		return false
	}
	for i := 1; i < len(intlist); i++ {
		diff := abs(intlist[i] - intlist[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
func isSafeAbs(intlist []int) bool {
	if len(intlist) < 1 {
		return false
	}
	increasing := true
	decreasing := true
	for i := 1; i < len(intlist); i++ {
		diff := intlist[i] - intlist[i-1]
		absdiff := abs(diff)
		if absdiff < 1 || absdiff > 3 {
			return false
		}
		if diff <= 0 {
			increasing = false
		}
		if diff >= 0 {
			decreasing = false
		}
	}
	return decreasing || increasing
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file ")
	}
	scanner := bufio.NewScanner(file)

	defer file.Close()
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		cuts := strings.Split(line, " ")

		intlist := []int{}
		for _, item := range cuts {
			integer, err := strconv.Atoi(item)
			if err != nil {
				fmt.Errorf("Error converting '%s' to integer '%v' \n", item, err)
				continue
			}
			intlist = append(intlist, integer)
		}
		if isSafeAbs(intlist) {
			count++
			continue
		}
		for i := 0; i < len(intlist); i++ {
			sliced_list := append([]int{}, intlist[:i]...)
			sliced_list = append(sliced_list, intlist[i+1:]...)
			fmt.Printf("In Option %x\n", sliced_list)
			if isSafeAbs(sliced_list) {
				fmt.Printf("Safe %x - %x \n", intlist, sliced_list)
				count++
				break
			}
		}

	}
	fmt.Println(count)
}
