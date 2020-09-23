package cryptopro

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgduncan/CryptoPro-Alexa-GCP/shared"
)

var coinbaseClient *shared.CoinbaseClient

func init() {
	client := http.DefaultClient
	coinbaseClient = &shared.CoinbaseClient{
		HTTP: client,
	}

}

// FetchPriceHTTP http entrypoint for cloud function
func FetchPriceHTTP(w http.ResponseWriter, r *http.Request) {

	resp, err := coinbaseClient.GetSpotPrice(r.Context())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
