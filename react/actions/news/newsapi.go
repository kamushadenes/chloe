package news

import (
	"encoding/json"
	"fmt"
	"github.com/biter777/countries"
	"github.com/kamushadenes/chloe/config"
	"net/http"
	"strings"
)

func NewsAPIQuery(q string) (*NewsAPIResponse, error) {
	var newsResp NewsAPIResponse

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/everything", nil)
	if err != nil {
		return nil, err
	}

	values := req.URL.Query()
	values.Add("q", q)
	values.Add("sortBy", config.React.NewsAPISortStrategy)
	values.Add("apiKey", config.React.NewsAPIToken)

	req.URL.RawQuery = values.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&newsResp); err != nil {
		return nil, err
	}

	return &newsResp, nil
}

func convertCountry(c string) (string, error) {
	country := countries.ByName(c)

	if !country.IsValid() {
		return "", fmt.Errorf("invalid country: %s", c)
	}

	return strings.ToLower(country.Alpha2()), nil
}

func NewsAPITopHeadlines(country string) (*NewsAPIResponse, error) {
	var newsResp NewsAPIResponse

	req, err := http.NewRequest("GET", "https://newsapi.org/v2/top-headlines", nil)
	if err != nil {
		return nil, err
	}

	vCountry, err := convertCountry(country)
	if err != nil {
		return nil, err
	}

	values := req.URL.Query()
	values.Add("country", vCountry)
	values.Add("apiKey", config.React.NewsAPIToken)

	req.URL.RawQuery = values.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&newsResp); err != nil {
		return nil, err
	}

	return &newsResp, nil
}
