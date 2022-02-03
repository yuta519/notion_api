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

func FetchfDatabase(secret_token string, db_id string) utils.Object {
	response := http.Get(utils.BaseUrl+"databases/"+db_id, secret_token)
	var database utils.Object
	err := json.Unmarshal(response, &database)
	if err != nil {
		os.Exit(1)
	}
	return database
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

// TODO: could make?
// func CreateDatabase(
// 	secret_token string,
// 	parent_page_id string,
// 	// title []map[string]interface{},
// 	key string,
// 	attribute string,
// 	value string,
// ) {
// 	payload := "{\"parent\": {\"type\": \"page_id\", \"page_id\": \"" + parent_page_id + "\"}, " +
// 		"\"properties\": {\"" + key + "\": {\"" + attribute + "\": [{\"text\": {\"content\": \"" + value + "\"}}]}}}"
// 	res := http.Post(utils.BaseUrl+"databases", secret_token, payload)
// 	fmt.Println(string(res))
// }

func ExportDbToMarkdown(secret_token string, db_id string) {
	database := FetchfDatabase(secret_token, db_id)
	title_of_db := database.Title
	db_name := "Untitled"

	if len(title_of_db) > 0 {
		db_name = title_of_db[0].PlainText
	}

	// Pass the error handling
	os.Mkdir(db_name, 0777)
	// if err := os.Mkdir(db_name, 0777); err != nil {
	// 	os.Exit(1)
	// }

	// Pass the database markdown
	// file, err := os.Create("./" + db_name + "/" + db_name + ".md")
	// if err != nil {
	// 	os.Exit(1)
	// }
	// file.WriteString("# " + db_name + "\n\n\n")
	// pages := FetchPagesByDbId(secret_token, db_id)
	// for property := range database.Properties {
	// 	for _, page := range pages.Results {
	// 		page_type := page.Properties[property].(map[string]interface{})["type"].(string)
	// 		if page_type == "rich_text" {
	// 			rich_text := page.Properties[property].(map[string]interface{})["rich_text"].([]interface{})
	// 			if len(rich_text) > 0 {
	// 				file.WriteString(property + ": " +
	// 					rich_text[0].(map[string]interface{})["plain_text"].(string),
	// 				)
	// 			}
	// 		} else if page_type == "multi_select" {
	// 			fmt.Println()
	// 		}
	// 		file.WriteString("\n")
	// 	}
	// }

	pages := FetchPagesByDbId(secret_token, db_id)
	for _, page := range pages.Results {
		ExportPageToMarkdown(secret_token, page.Id, db_name)
	}
}
