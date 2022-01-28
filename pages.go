package notion_api

import (
	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchPageList(secret_token string, db_id string) string {
	response := http.Post(
		utils.BaseUrl+"databases/"+db_id+"/query",
		secret_token,
		"{\"page_size\":100}",
	)
	return response
}

func CreatePage(secret_token string, db_id string, content string) string {
	response := http.Post(
		utils.BaseUrl+"pages",
		secret_token,
		"{\"parent\": {\"database_id\": \""+db_id+"\"}, "+
			// TODO: change `Name` to be able to input from argument
			"\"properties\": {\"Name\": {\"title\": "+
			"[{\"text\": {\"content\": \""+content+"\"}}]}}}",
	)
	return response
}

// TODO: create schema.go to define attribute
func UpdatePage(
	secret_token string,
	page_id string,
	where string,
	attribute string,
	content string,
) string {
	response := http.Patch(
		utils.BaseUrl+"pages/"+page_id,
		secret_token,
		"{\"properties\": { \""+where+"\": {\""+attribute+
			"\": [{\"text\": {\"content\": \""+content+"\"}}]}}}",
	)
	return response
}
