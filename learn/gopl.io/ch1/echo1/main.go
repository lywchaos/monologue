package main

import (
	"fmt"
	"os"
)

func main() {
	for _, item := range(os.Args[1:]) {
		fmt.Printf("%v ", item)
	}
}