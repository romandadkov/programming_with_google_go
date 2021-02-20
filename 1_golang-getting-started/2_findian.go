package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main(){

    var s string

    fmt.Printf("enter a string: ")
    input := bufio.NewReader(os.Stdin)
    raw_line, err := input.ReadString('\n')

    if (err == nil){
        s = strings.TrimSuffix(raw_line, "\n")
        s = strings.ToLower(s)

        pref := strings.HasPrefix(s, "i")
        suff := strings.HasSuffix(s, "n")
        contains := strings.Contains(s, "a")

        if (pref && contains && suff){
            fmt.Printf("Found!\n")
        } else{
            fmt.Printf("Not Found!\n")
        }
    } else{
        fmt.Printf("Input Error\n")
    }
}
