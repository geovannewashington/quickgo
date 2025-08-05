package main

import (
	"fmt"
	"myownlib/mylib"
)

func main() {
	result := mylib.Add(10, 5)
	fmt.Println("Result:", result)

	// mylib.subtract(10, 5) -> will not compile because subtract is not exported
}
