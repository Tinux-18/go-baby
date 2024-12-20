package main

import (
	"fmt"

	"github.com/Tinux-18/go-baby/internal/wiki"
)

func main() {
	// Get wiki output.
	wikiData := wiki.Get()
	Create(wikiData)
	fmt.Println(wikiData)
}
