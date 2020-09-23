package shared

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	cburl = "https://api.coinbase.com/v2/prices/spot?currency=USD"
)

// CoinbaseClient test
type CoinbaseClient struct {
	HTTP *http.Client
}

// CoinbaseResponse test
type CoinbaseResponse struct {
	Data coinbaseData `json:"data"`
}

type coinbaseData struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// GetSpotPrice test
func (c *CoinbaseClient) GetSpotPrice(ctx context.Context) (CoinbaseResponse, error) {
	url, err := url.Parse(cburl)
	if err != nil {
		return CoinbaseResponse{}, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), http.NoBody)
	if err != nil {
		return CoinbaseResponse{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req.WithContext(ctx))
	if err != nil {
		return CoinbaseResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return CoinbaseResponse{}, errors.New("Incorrect response body")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CoinbaseResponse{}, err
	}

	var cbResp CoinbaseResponse
	if err := json.Unmarshal(body, &cbResp); err != nil {
		return CoinbaseResponse{}, err
	}

	return cbResp, nil
}
