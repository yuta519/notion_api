package notion_api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

type Databases struct {
	HasMore    bool                     `json:"has_more"`
	NextCursor interface{}              `json:"next_cursor"`
	Object     string                   `json:"object"`
	Results    []map[string]interface{} `json:"results"`
}

type Database struct {
	Cover          *Cover `json:"cover"`
	CreatedTime    string `json:"created_time"`
	Icon           *Icon  `json:"icon"`
	Id             string `json:"id"`
	LastEditedTime string `json:"last_edited_time"`
}

type Cover struct {
}

type Icon struct {
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

func FetchDatabases(secret_token string) {
	response := FetchRawResponseOfDatabases(secret_token).Results
	fmt.Println(response[0])
}
