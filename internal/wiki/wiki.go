package wiki

import (
	"encoding/json"
	"log"
	"strings"

	"cgt.name/pkg/go-mwclient"
	"github.com/antonholmquist/jason"
)

type WikiResponse struct {
	Batchcomplete bool `json:"batchcomplete"`
	Query         struct {
		Categorymembers []WikiName `json:"categorymembers"`
	} `json:"query"`
}

type WikiName struct {
	PageID  int    `json:"pageid"`
	NS      int    `json:"ns"`
	Title   string `json:"title"`
	PageURL string
}

var client *mwclient.Client

func init() {
	resp, err := mwclient.New("https://en.wikipedia.org/w/api.php", "RomanianBabyBoyNames/1.0")
	client = resp

	if err != nil {
		log.Fatal(err)
	}
}

func processResp(resp *jason.Object) []WikiName {
	var wikiResp WikiResponse
	err := json.Unmarshal([]byte(resp.String()), &wikiResp)
	if err != nil {
		log.Fatal(err)
	}
	return wikiResp.Query.Categorymembers
}

func getURL(title string) string {
	return "https://en.wikipedia.org/wiki/" + title
}

func cleanName(name string) string {
	idx := strings.Index(name, "(")

	if idx == -1 {
		return name
	}

	return name[:idx-1]
}

func prettify(names []WikiName) []WikiName {
	for i := range names {
		names[i].PageURL = getURL(strings.ReplaceAll(names[i].Title, " ", "_"))
		names[i].Title = cleanName(names[i].Title)
	}
	return names
}

func Get() []WikiName {

	// Set up the query parameters.
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

	return prettify(processResp(resp))
}
