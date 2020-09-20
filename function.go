package cryptopro

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	coinbaseBTCURL = "https://api.coinbase.com/v2/prices/spot?currency=USD"
)

type coinbaseResponse struct {
	Data coinbaseData `json:"data"`
}

type coinbaseData struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// FetchPriceHTTP http entrypoint for cloud function
func FetchPriceHTTP(w http.ResponseWriter, r *http.Request) {
	client := http.DefaultClient

	resp, err := client.Get(coinbaseBTCURL)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var pd coinbaseResponse
	json.Unmarshal(body, &pd)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pd)
}
