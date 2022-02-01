package notion_api

import (
	"encoding/json"
	"fmt"
	"os"
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

func ExportPageToMarkdown(secret_token string, page_id string) {
	blocks := FetchChildrenInBlock(secret_token, page_id, 1000)

	file, err := os.Create("./file.md")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	for _, block := range blocks.Results {
		if block.Type == "paragraph" {
			file.WriteString(parseParagraphToMarkdown(block))
		} else if block.Type == "heading_1" {
			file.WriteString(parseHeading1ToMarkdown(block))
		} else if block.Type == "heading_2" {
			file.WriteString(parseHeading2ToMarkdown(block))
		} else if block.Type == "heading_3" {
			file.WriteString(parseHeading3ToMarkdown(block))
		} else if block.Type == "to_do" {
			file.WriteString(parseToDoToMarkdown(block))
		} else if block.Type == "child_page" {
			file.WriteString(parseChildPageToMarkdown(block))
		} else if block.Type == "child_database" {
			fmt.Println(block.Type)
			fmt.Println(block)
		} else if block.Type == "table" {
			// can't export table type because api does not support yet.
			file.WriteString(parseTableToMarkdown(block))
		} else if block.Type == "bulleted_list_item" {
			file.WriteString(parseBulletedListItemToMarkdown(block))
		} else if block.Type == "numbered_list_item" {
			// can't export table type because api does not support yet.
			file.WriteString(parseNumberedListItemToMarkdown(block))
		} else if block.Type == "toggle" {
			// can't export table type because api does not support yet.
			file.WriteString(parseToggleToMarkdown(block))
		} else if block.Type == "quote" {
			file.WriteString(parseQuoteToMarkdown(block))
		} else if block.Type == "divider" {
			file.WriteString(parseDividerToMarkdown(block))
		} else if block.Type == "callout" {
			file.WriteString(parseCalloutToMarkdown(block))
		} else {
			fmt.Println(block.Type)
			fmt.Println(block)
		}
		file.WriteString("\n")
	}
}
