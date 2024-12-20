package json

import (
	"encoding/json"
	"github.com/Tinux-18/go-baby/internal/wiki"
	"log"
	"os"
)

type BabyName struct {
	Name    string `json:"title"`
	PageURL string
}

type Filter struct {
	Names []BabyName
	Favs  []BabyName
}

func getNames(wikiData []wiki.WikiName) []BabyName {
	babyNames := make([]BabyName, len(wikiData))
	for i, name := range wikiData {
		babyNames[i] = BabyName{
			Name:    name.Title,
			PageURL: name.PageURL,
		}
	}
	return babyNames
}

func Create(wikiData []wiki.WikiName) {
	data := Filter{Names: getNames(wikiData), Favs: []BabyName{}}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("names.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
