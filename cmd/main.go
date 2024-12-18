package main

import (
	"fmt"
	"go-baby/wiki"
)

type BabyName struct {
	name string
	page string // coudl use client.GetPagesByID to get the url.
}

func main() {
	// Get wiki output.
	jason := wiki.Get()

	fmt.Println(jason)
}
