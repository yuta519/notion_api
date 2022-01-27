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
