package cryptopro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type coinbaseResponse struct {
	Data coinbaseData `json:"data"`
}

type coinbaseData struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"Amount"`
}

// FetchPriceHTTP http entrypoint for cloud function
func FetchPriceHTTP(w http.ResponseWriter, r *http.Request) {
	client := http.DefaultClient

	resp, err := client.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var priceData coinbaseResponse
	json.Unmarshal(body, &priceData)

	fmt.Fprintln(w, priceData.Data.Amount)
}
