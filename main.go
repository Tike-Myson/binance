package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Trade []struct {
	ID           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func main() {
	for {
		time.Sleep(5 * time.Second)
		TradeAPI()
	}
}

func TradeAPI()  {
	url := "https://api1.binance.com/api/v3/trades?symbol=BTCUSDT&limit=2"
	trade := new(Trade)
	resp, err :=  http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	json.Unmarshal(body, &trade)

	fmt.Println(trade)

}
