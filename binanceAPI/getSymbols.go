package binanceAPI

import (
	"encoding/json"
	"log"
	"net/http"

	config "github.com/bot/config"
	structures "github.com/bot/struct"
)

func GetSymbols() []string {
	resp, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")

	if err != nil {
		log.Println("err get exchangeInfo: ", err)
	}

	defer resp.Body.Close()

	var dataSt structures.ExchInfo

	err = json.NewDecoder(resp.Body).Decode(&dataSt)

	if err != nil {
		log.Println("err get exchangeInfo: ", err)
	}
	symbLengt := len(dataSt.Symbols)
	symbols := make([]string, symbLengt)

	for i, symb := range dataSt.Symbols {
		symbols[i] = symb.Baseasset
	}

	rd := RemoveDuplic(symbols)
	return RemoveUnusedSymb(rd)
}

func RemoveUnusedSymb(s []string) []string {
	for i := 0; i < len(s); i++ {
		url := s[i]
		for _, rem := range config.RemoveSymb {
			if url == rem {
				s = append(s[:i], s[i+1:]...)
				i--
				break
			}
		}
	}
	return s
}

func RemoveDuplic(s []string) []string {
	m := make(map[string]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
		} else {
			m[item] = true
		}
	}

	var result []string
	for item := range m {
		result = append(result, item)
	}
	return result
}
