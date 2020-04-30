package main

import (
	"fmt"
	"os"

	optimizer ".."
)

// CatchPanic recovers panics from inner stack
func CatchPanic() {
	if a := recover(); a != nil {
		fmt.Println(a)
		os.Exit(0)
	}
}

func main() {
	defer CatchPanic()
	args := os.Args[1:]
	if len(args) != 1 {
		panic("One argument required")
	}
	filename := args[0]
	//filename := "../sample/goodexample03.txt"
	data := optimizer.New(filename) //.Solve()
	result := data.Solve()
	optimizer.Print(result)
}
