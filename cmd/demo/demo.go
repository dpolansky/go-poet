package main

import "fmt"

func main() {
	fmt.Println("Calling hello...")
	hello()
}

func hello() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
