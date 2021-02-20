package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the filename:")
	scanner.Scan()
	fileName := scanner.Text()
	var handler *os.File
	persons := make([]Person, 0, 1)

	for {
		fin, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			fmt.Println("Please Enter valid file name:")
			scanner.Scan()
			fileName = scanner.Text()
		} else {
			handler = fin
			break
		}
	}

	reader := bufio.NewScanner(handler)
	var arr []string

	for reader.Scan() {
		arr = strings.Split(reader.Text(), " ")
		persons = append(persons, Person{arr[0], arr[1]})
	}

	handler.Close()

	for _, person := range persons {
		fmt.Printf("%v - %v\n", person.fname, person.lname)
	}
}
