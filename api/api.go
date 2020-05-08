package api

import (
	"inabajunmr/dicc/config"
	"inabajunmr/dicc/dictionary"
)

type ApiType int

const (
	WEBSTER = iota
)

type DictionaryApi interface {
	SearchWords(SearchCondition) (dictionary.Result, error)
}

type SearchCondition struct {
	Word string
}

func GetApi(apiType ApiType) DictionaryApi {
	switch apiType {
	case WEBSTER:
		key := config.GetMerriamWebsterApiKey()
		return Webster{key}
	default:
		return nil
	}

}
