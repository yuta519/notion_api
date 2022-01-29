package notion_api

import (
	"strconv"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchPages(
	secret_token string,
	db_id string,
	page_size int,
) []byte {
	response := http.Post(
		utils.BaseUrl+"databases/"+db_id+"/query",
		secret_token,
		"{\"page_size\":"+strconv.Itoa(page_size)+"}",
	)
	return response
}

func CreatePage(
	secret_token string,
	db_id string,
	key string,
	attribute string,
	value string,
) []byte {
	response := http.Post(
		utils.BaseUrl+"pages",
		secret_token,
		"{\"parent\": {\"database_id\": \""+db_id+"\"}, "+
			"\"properties\": {\""+key+"\": {\""+attribute+"\": "+
			"[{\"text\": {\"content\": \""+value+"\"}}]}}}",
	)
	return response
}

func UpdatePage(
	secret_token string,
	page_id string,
	key string,
	attribute string,
	value string,
) []byte {
	response := http.Patch(
		utils.BaseUrl+"pages/"+page_id,
		secret_token,
		"{\"properties\": { \""+key+"\": {\""+attribute+
			"\": [{\"text\": {\"content\": \""+value+"\"}}]}}}",
	)
	return response
}
