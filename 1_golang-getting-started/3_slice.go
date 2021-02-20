package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int = make([]int, 3)
	var in string

	for true {
		fmt.Println("Please again enter an interger (X to exit):")
		fmt.Scanln(&in)
		if in == "X" {
			break
		}
		n, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Wrong input")
		} else {
			s = append(s, n)
			sort.Ints(s[:])
			fmt.Println(s)
		}
	}
}

