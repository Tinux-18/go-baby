package main

import (
	"fmt"
	"go-baby/wiki"
)

func main() {
	// Get wiki output.
	jason := wiki.Get()

	fmt.Println(jason)
}
