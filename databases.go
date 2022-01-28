package notion_api

import (
	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchDatabaseIds(secret_token string) map[string]interface{} {
	response := http.Get(utils.BaseUrl+"databases", secret_token)
	return response
}
