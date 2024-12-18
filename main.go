package main

import (
	"fmt"
	"log"

	"cgt.name/pkg/go-mwclient"
)

func main() {
	// Initialize the client
	client, err := mwclient.New("https://en.wikipedia.org/w/api.php", "RomanianBabyBoyNames/1.0")
	if err != nil {
		log.Fatal(err)
	}

	// Set up the query parameters
	params := map[string]string{
		"action":     "query",
		"list":       "categorymembers",
		"cmtitle":    "Category:Romanian_masculine_given_names",
		"cmlimit":    "500",
		"cmcontinue": "",
	}

	// Make the request.
	resp, err := client.Get(params)
	if err != nil {
		panic(err)
	}

	// Print the *jason.Object
	fmt.Println(resp)
}
