package main

import (
	"fmt"
	"wait2die"
)

func main() {
	fmt.Println("Hello World")
	wait2die.WaitToDie(func() {
		fmt.Println("Before quit")
	})
}