package json

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Tinux-18/go-baby/internal/wiki"
)

type BabyName struct {
	Name    string
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

func Create(wikiData []wiki.WikiName, filename string) {
	// Create JSON data.
	data := Filter{Names: getNames(wikiData), Favs: []BabyName{}}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Create file.
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}
