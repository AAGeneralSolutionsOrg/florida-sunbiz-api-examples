package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("RAPIDAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Set RAPIDAPI_KEY before running this example")
		os.Exit(1)
	}

	const host = "florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com"
	endpoint := url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     "/v1/entity/search",
		RawQuery: url.Values{"name": {"publix"}, "limit": {"2"}}.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-RapidAPI-Key", apiKey)
	req.Header.Set("X-RapidAPI-Host", host)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	var body any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}

	formatted, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(formatted))
}
