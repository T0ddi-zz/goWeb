package main

import (
	"fmt"
	"os"
)

func main() {
	// UNCOMMENT & FIX THIS CODE
	//count := len(os.Args) - 1

	l := len(os.Args) - 1
	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Println("There are", l, "peolpe!")
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Hello great", os.Args[4], "!")
	fmt.Println("Hello great", os.Args[5], "!")
	//  Nice to meet you all

	fmt.Println("Nice to meet you all")
}
