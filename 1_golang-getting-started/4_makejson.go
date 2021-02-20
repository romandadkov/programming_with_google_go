package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)
	var name string
	var addr string
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter your address")
	fmt.Scanln(&addr)
	person["name"] = name
	person["address"] = addr
	json, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

