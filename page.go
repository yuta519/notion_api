package notion_api

import (
	"github.com/yuta519/notion_api/lib"
)

var (
	base_url string = "https://api.notion.com/v1/databases/"
)

func fetch_page_id(secret_token string, db_id string) string {
	response := lib.PostRequest(base_url+db_id+"/query", secret_token, "{\"page_size\":100}")
	return response
}
