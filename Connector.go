package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"./calculator"
)

func ReadBitBay() {
	fetch(calculator.BTC, "PLN")
	fetch(calculator.ETH, "PLN")
	fetch(calculator.LTC, "PLN")
	fetch(calculator.LSK, "PLN")
	fetch(calculator.GAME, "PLN")
	fetch(calculator.BCC, "PLN")
	fetch(calculator.DASH, "PLN")
}

var calc = calculator.New()

func fetch(crypto calculator.CryptoType, currency string) {
	uri := fmt.Sprintf("https://bitbay.net/API/Public/%s%s/ticker.json", crypto, currency)
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	dict := make(map[string]interface{})
	err = json.Unmarshal(body, &dict)
	if err != nil {
		panic(err)
	}
	value := dict["last"].(float64)

	calc.SetValue(crypto, value)
}
