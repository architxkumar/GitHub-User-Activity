package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid number of arguments")
	}
	username := arguments[1]
	fmt.Printf("Hello, %s!\n", username)
}
