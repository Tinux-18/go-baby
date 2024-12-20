package main

import (
	"fmt"
	"go-baby/internal/wiki"
)

func main() {
	// Get wiki output.
	wikiData := wiki.Get()

	fmt.Println(wikiData)
}
