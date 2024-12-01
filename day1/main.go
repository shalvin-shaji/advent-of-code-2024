package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(number int) int{
    if number >= 0{
        return number
    }
    return -number
}

func main() {
    list1 := []int {}
    list2 := []int {}
    file, err := os.Open("day1.txt")
    if err != nil{
        fmt.Printf("Error reading the file: %e", err)
        return
    }
    defer file.Close()

    scanner:= bufio.NewScanner(file)
    distance := 0
    for scanner.Scan() {
        line := scanner.Text()
        n1, n2, found := strings.Cut(line, " ")
        if !found {
            fmt.Println("could not cut the string into two")
        }
        number1, err := strconv.Atoi(strings.TrimSpace(n1))
        if err != nil{
            fmt.Printf("Error converting first string to integer: %e", err)
            return
        }
        number2, err := strconv.Atoi(strings.TrimSpace(n2))
        if err != nil{
            fmt.Printf("Error converting second string to integer: %e", err)
            return
        }
        list1 = append(list1, number1)
        list2 = append(list2, number2)

    }
        sort.Ints(list1)
        sort.Ints(list2)
        fmt.Println(list1)
        fmt.Println(list2)
        for i := range list1{
            distance += abs(list1[i] - list2[i])
        }
    fmt.Println(distance)
}


