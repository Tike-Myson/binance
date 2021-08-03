package main

<<<<<<< HEAD
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Trade struct {
	ID           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func main() {
	fmt.Printf("|%20s|%20s|%20s|\n", "Price", "Qty", "QuoteQty")
	for {
		time.Sleep(1 * time.Second)
		TradeAPI()
	}
}

func TradeAPI()  {
	url := "https://api1.binance.com/api/v3/trades?symbol=BTCUSDT&limit=2"
	var trades []Trade
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

	json.Unmarshal(body, &trades)

	for _, v := range trades {
		fmt.Printf("|%20s|%20s|%20s|\n", v.Price, v.Qty, v.QuoteQty)
	}
=======
import "fmt"

func main() {
	fmt.Println("Hello World")
>>>>>>> 0c9fdd9089bc5d486f9d6ff9919062e246a0f0c4
}
