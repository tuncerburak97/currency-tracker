package currency

import (
	"currency-tracker/internal/config"
	client "currency-tracker/internal/http"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	config     *config.Config
}

var (
	clientInstance *Client
)

func GetCurrencyClient() *Client {
	if clientInstance == nil {
		clientInstance = &Client{
			httpClient: client.GetHttpClient(),
			config:     config.GetConfig(),
		}
	}
	return clientInstance
}

func (c *Client) FetchGoldPrices() ([]GoldResponse, error) {
	url := c.config.Rest.Altinkaynak.Gold

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var goldPrices []GoldResponse
	if err := json.NewDecoder(resp.Body).Decode(&goldPrices); err != nil {
		return nil, err
	}

	return goldPrices, nil
}

func (c *Client) FetchCurrencyPrices() ([]GetCurrencyResponse, error) {
	url := c.config.Rest.Altinkaynak.Currency

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var currencyPrices []GetCurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyPrices); err != nil {
		return nil, err
	}

	return currencyPrices, nil
}
