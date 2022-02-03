package notion_api

import (
	"encoding/json"
	"fmt"
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

	// if err := os.Mkdir(db_name, 0777); err != nil {
	// 	os.Exit(1)
	// }

	// file, err := os.Create("./" + db_name + "/" + db_name + ".md")
	// if err != nil {
	// 	os.Exit(1)
	// }

	// fmt.Println(database.ObjectType)
	fmt.Println(db_name)

	pages := FetchPagesByDbId(
		"secret_OgKvRWjGQu3fzSHuApNIkFXKu4nxjiw3TOahfguoIPA",
		"17cb9b38-1749-46f0-9b86-8c2b77abd898",
	)

	for property := range database.Properties {
		fmt.Println(property)
		fmt.Println()
		for _, page := range pages.Results {
			fmt.Println(page.Properties[property])
			fmt.Println()
		}
	}

	// file.WriteString("# " + db_name)

	for _, page := range pages.Results {
		fmt.Println(page.Id)
		fmt.Println(page.Properties)
		fmt.Println()
	}

}
