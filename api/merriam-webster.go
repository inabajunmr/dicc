package api

import (
	"fmt"
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

func (w Webster) SearchWords(cond SearchCondition) ([]dictionary.Word, error) {

	response, err := http.Get(assembleUrl(w.apiKey, cond))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return nil, nil
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
