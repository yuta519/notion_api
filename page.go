package notion_api

import (
	"github.com/yuta519/notion_api/handler/http"
)

var (
	base_url string = "https://api.notion.com/v1/"
)

func FetchPageList(secret_token string, db_id string) string {
	response := http.PostRequest(
		base_url+"databases/"+db_id+"/query",
		secret_token,
		"{\"page_size\":100}",
	)
	return response
}

func CreatePage(secret_token string, db_id string, content string) string {
	response := http.PostRequest(
		base_url+"pages",
		secret_token,
		"{\"parent\": {\"database_id\": \""+db_id+"\"}, "+
			// TODO: change `Name` to be able to input from argument
			"\"properties\": {\"Name\": {\"title\": "+
			"[{\"text\": {\"content\": \""+content+"\"}}]}}}",
	)
	return response
}
