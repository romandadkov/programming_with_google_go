package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_ints() []int {

	read_input := bufio.NewScanner(os.Stdin)
	fmt.Println("\nPlease enter a sequence of integers separated by spaces.")
	fmt.Println("Hit enter when finished:")
	read_input.Scan()
	line := read_input.Text()

	user_ints := strings.Split(line, " ")

	var user_slice []int

	for i := 0; i < len(user_ints); i++ {
		user_int, err := strconv.Atoi(user_ints[i])
		if err == nil {
			user_slice = append(user_slice, user_int)
		} else {
			fmt.Println("Invalid input. Please run the program again with ")
			fmt.Println("a sequence of integers separated by spaces. ")
			fmt.Print("Exiting program ")
			os.Exit(1)
		}

	}
	return user_slice

}

func main() {
	fmt.Println("\nSorting partitions ...")
	arr := read_ints()
	
	// split
	quart1 := len(arr) / 4
	quart2 := len(arr) / 2
	quart3 := quart1 + quart2
	
	// sort partitions
	c := make(chan []int, 4)
	go sort_array(arr[:quart1], c)
	go sort_array(arr[quart1:quart2], c)
	go sort_array(arr[quart2:quart3], c)
	go sort_array(arr[quart3:], c)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c
	
	// merge
	final := merge(merge(arr1, arr2), merge(arr3, arr4))
	fmt.Println(final)
}

func sort_array(arr []int, c chan []int) {
	orig_arr := make([]int, len(arr))
	copy(orig_arr, arr)
	sort.Ints(arr)
	fmt.Println("Sorted", orig_arr, "into", arr)

	c <- arr
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
