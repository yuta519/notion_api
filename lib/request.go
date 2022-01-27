package lib

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	notion_version string = "2021-08-16"
	content_type   string = "application/json"
)

func GetRequest(endpoint string, secret_token string) string {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Authorization", secret_token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func PostRequest(endpoint string, secret_token string, payload string) string {
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(payload))
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Authorization", secret_token)
	req.Header.Set("Content-Type", content_type)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
