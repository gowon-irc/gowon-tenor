package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

const (
	tenorAPIURLbase = "https://api.tenor.com/v1/search"
)

func searchURL(term string, apiKey string) (url string) {
	temp := "%s?q=%s&key=%s&limit=10"

	return fmt.Sprintf(temp, tenorAPIURLbase, term, apiKey)
}

type tenorJSON struct {
	Results []struct {
		Media []struct {
			Gif struct {
				URL string `json:"url"`
			} `json:"gif"`
		} `json:"media"`
		Title string `json:"title"`
	} `json:"results"`
}

func tenor(args, apiKey string) (msg string, err error) {
	data := &tenorJSON{}
	msg = url.QueryEscape(args)

	url := searchURL(msg, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, data)

	if err != nil {
		return "", err
	}

	if len(data.Results) == 0 {
		return fmt.Sprintf("No %s gifs found", msg), nil
	}

	index := rand.Intn(len(data.Results))

	return data.Results[index].Media[0].Gif.URL, nil
}
