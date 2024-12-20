package main

import (
	"github.com/Tinux-18/go-baby/internal/json"
	"github.com/Tinux-18/go-baby/internal/wiki"
)

func main() {
	filename := "names.json"

	// Get wiki output.
	if !json.FileExists(filename) {
		wikiData := wiki.Get()
		json.Create(wikiData, filename)
	}
}
