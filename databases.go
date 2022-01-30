package notion_api

import (
	"encoding/json"
	"os"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

type Databases struct {
	HasMore    bool        `json:"has_more"`
	NextCursor interface{} `json:"next_cursor"`
	Object     string      `json:"object"`
	Results    []Database  `json:"results"`
}

type Database struct {
	Cover          *Cover                 `json:"cover"`
	CreatedTime    string                 `json:"created_time"`
	Icon           *Icon                  `json:"icon"`
	Id             string                 `json:"id"`
	LastEditedTime string                 `json:"last_edited_time"`
	Object         string                 `json:"object"`
	Parent         Parent                 `json:"parent"`
	Properties     map[string]interface{} `json:"properties"`
	Title          []Title                `json:"title"`
	Url            string                 `json:"url"`
}

type Cover struct {
	External map[string]string `json:"external"`
	Type     string            `json:"type"`
}

type Icon struct {
	Emoji string `json:"emoji"`
	Type  string `json:"type"`
}

type Parent struct {
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
}

type Title struct {
	Annotations map[string]interface{} `json:"annotations"`
	Href        *string                `json:"href"`
	PlainText   string                 `json:"plain_text"`
	Text        map[string]string      `json:"text"`
	Type        string                 `json:"type"`
}

func FetchRawResponseOfDatabases(secret_token string) Databases {
	response := http.Get(utils.BaseUrl+"databases", secret_token)
	var databases Databases
	err := json.Unmarshal(response, &databases)
	if err != nil {
		os.Exit(1)
	}
	return databases
}

func FetchDatabases(secret_token string) []Database {
	response := FetchRawResponseOfDatabases(secret_token).Results
	return response
}

func FetchDatabaseIds(secret_token string) []map[string]string {
	response := FetchRawResponseOfDatabases(secret_token).Results
	var database_ids []map[string]string
	for _, database := range response {
		database_ids = append(
			database_ids,
			map[string]string{
				"id":    database.Id,
				"title": database.Title[0].PlainText,
			},
		)
	}
	return database_ids
}
