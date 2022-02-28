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

func FetchPageByPageId(
	secret_token string,
	page_id string,
) utils.Object {
	response := http.Get(
		utils.BaseUrl+"pages/"+page_id,
		secret_token,
	)
	var page utils.Object
	json.Unmarshal(response, &page)
	return page
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

func ExportPageToMarkdown(
	secret_token string,
	page_id string,
	destination_directory string,
) {
	page_info := FetchPageByPageId(secret_token, page_id)
	blocks := FetchChildrenInBlock(secret_token, page_id, 1000)

	title_of_page_info := page_info.Properties["Name"].(map[string]interface{})["title"].([]interface{})
	filename := "Untitled"
	if len(title_of_page_info) > 0 {
		filename = title_of_page_info[0].(map[string]interface{})["plain_text"].(string)
	}

	file, err := os.Create(
		destination_directory + "/" + strings.ReplaceAll(filename, " ", "_") + ".md",
	)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString("# " + filename)
	file.WriteString("\n---\n\n\n")

	for i, block := range blocks.Results {
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
			// can't export child database type because does not support yet.
			fmt.Println(block.Type)
			fmt.Println(block)
		} else if block.Type == "table" {
			// can't export table type because api does not support yet.
			file.WriteString(parseTableToMarkdown(secret_token, block))
		} else if block.Type == "bulleted_list_item" {
			file.WriteString(parseBulletedListItemToMarkdown(block))
		} else if block.Type == "numbered_list_item" {
			number := 1
			if i != 0 {
				for i := i; blocks.Results[i-1].Type == "numbered_list_item"; i-- {
					number += 1
				}
			}
			file.WriteString(parseNumberedListItemToMarkdown(block, number))
		} else if block.Type == "toggle" {
			// can't export toggle type because api does not support yet.
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
