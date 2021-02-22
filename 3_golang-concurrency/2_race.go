package main

import (
	"fmt"
	"time"
)

func printer(a string) {

	fmt.Println(a)

}

func main() {

	for range [5]int{} {
		go printer("first")
		go printer("second")
		time.Sleep(10 * time.Millisecond)
	}
}

/*
The program has a simple function that prints a string.
This is called from within a loop as 2 goroutines:
go printer("first")
go printer("second")
"first" and "second" printed by the goroutines appear at
unreliable intervals, thus containing a race condition.
As you can see, the output is different each time the program is run,
indicating the race condition.
$ go run race.go
first
second
second
first
second
first
first
second
second
first
$ go run race.go
second
first
second
first
first
second
first
second
first
second
*/
