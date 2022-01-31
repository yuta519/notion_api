package notion_api

import (
	"encoding/json"
	"strings"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

type Properties struct {
	Attribute string
	Key       string
	Value     string
}

func FetchPagesByDbId(
	secret_token string,
	db_id string,
) utils.Objects {
	response := http.Post(
		utils.BaseUrl+"databases/"+db_id+"/query",
		secret_token,
		"",
	)
	var pages utils.Objects
	json.Unmarshal(response, &pages)
	return pages
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

func UpdatePropertiesInPage(
	secret_token string,
	page_id string,
	properties []Properties,
) string {
	payloads := "{\"properties\": {"
	for _, property := range properties {
		payload := "\"" + property.Key + "\": {\"" +
			property.Attribute + "\": [{\"text\": {\"content\": \"" +
			property.Value + "\"}}]},"
		payloads = payloads + payload
	}
	payloads = strings.TrimRight(payloads, ",")
	payloads = payloads + "}}"

	response := http.Patch(
		utils.BaseUrl+"pages/"+page_id,
		secret_token,
		payloads,
	)
	return string(response)
}
