package main

import "fmt"

func main(){

    var f float64

    fmt.Printf("enter floating point number: ")
    num, err := fmt.Scan(&f)
    t := int(f)

    if (num == 1 && err == nil){
        fmt.Printf("converted interger: %d\n", t)
    } else{
        fmt.Printf("input error\n")
    }
}
