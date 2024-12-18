package wiki

import (
	"encoding/json"
	"fmt"
	"log"

	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
)

type WikiResponse struct {
	Batchcomplete bool `json:"batchcomplete"`
	Query         struct {
		Categorymembers []BabyName `json:"categorymembers"`
	} `json:"query"`
}

type BabyName struct {
	PageID  int    `json:"pageid"`
	NS      int    `json:"ns"`
	Title   string `json:"title"`
	PageURL string
}

func processResponse(resp *jason.Object) []BabyName {
	var wikiResp WikiResponse
	err := json.Unmarshal([]byte(resp.String()), &wikiResp)
	if err != nil {
		log.Fatal(err)
	}
	return wikiResp.Query.Categorymembers
}

func getPageUrl(names []BabyName) []BabyName {
	client, err := mwclient.New("https://en.wikipedia.org/w/api.php", "RomanianBabyBoyNames/1.0")
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		resp, err := client.GetPageByID(name.PageID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
		name.PageURL = "https://en.wikipedia.org/wiki/"
	}
	return names
}

func Get() []BabyName { // Initialize the client
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
		log.Fatal(err)
	}

	return processResponse(resp)
}
