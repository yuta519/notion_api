package notion_api

import (
	"encoding/json"
	"os"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchRawResponseOfDatabases(secret_token string) utils.Objects {
	response := http.Get(utils.BaseUrl+"databases", secret_token)
	var databases utils.Objects
	err := json.Unmarshal(response, &databases)
	if err != nil {
		os.Exit(1)
	}
	return databases
}

func FetchDatabases(secret_token string) []utils.Object {
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
