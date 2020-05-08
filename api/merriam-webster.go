package api

import (
	"encoding/json"
	"inabajunmr/dicc/dictionary"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Webster struct {
	apiKey string
}

const URL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"

func (w Webster) SearchWords(cond SearchCondition) (dictionary.Result, error) {

	response, err := http.Get(assembleUrl(w.apiKey, cond))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// API Response of webster has not perfect matching word so need to parse word(when searching "test", return "test-flight"
	var websterResult []WebsterResult
	if err := json.Unmarshal(body, &websterResult); err != nil {
		log.Fatal(err)
	}

	var defs []dictionary.Definition
	for _, r := range websterResult {
		word := dictionary.Definition{r.Shortdefs, r.Fl}
		defs = append(defs, word)
	}

	return dictionary.Result{
		SearchWord:  cond.Word,
		Definitions: defs,
	}, nil
}

func assembleUrl(apiKey string, cond SearchCondition) string {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, cond.Word)
	q := u.Query()
	q.Set("key", apiKey)
	u.RawQuery = q.Encode()
	return u.String()
}

type WebsterResult struct {
	Fl        string   `json:"fl"`
	Shortdefs []string `json:"shortdef"`
}
