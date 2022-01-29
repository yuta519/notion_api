package http

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

func Get(endpoint string, secret_token string) []byte {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Authorization", secret_token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	responseByte, _ := ioutil.ReadAll(res.Body)
	return responseByte
}

func Post(
	endpoint string,
	secret_token string,
	payload string,
) []byte {
	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(payload))
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Authorization", secret_token)
	req.Header.Set("Content-Type", content_type)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	responseByte, _ := ioutil.ReadAll(res.Body)
	return responseByte

}

func Patch(
	endpoint string,
	secret_token string,
	payload string,
) []byte {
	req, _ := http.NewRequest("PATCH", endpoint, strings.NewReader(payload))
	req.Header.Set("Notion-Version", notion_version)
	req.Header.Set("Authorization", secret_token)
	req.Header.Set("Content-Type", content_type)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}

	defer res.Body.Close()
	responseByte, _ := ioutil.ReadAll(res.Body)
	return responseByte
}

// func parseResponseToJson(response string) map[string]interface{} {
// 	resBytes := []byte(response)
// 	var jsonRes map[string]interface{}
// 	err := json.Unmarshal(resBytes, &jsonRes)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	return jsonRes
// }
