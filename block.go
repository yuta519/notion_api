package notion_api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchBlock(secret_token string, block_id string) {
	response := http.Get(utils.BaseUrl+"blocks/"+block_id, secret_token)
	fmt.Println(string(response))
}

func FetchChildrenInBlock(
	secret_token string,
	block_id string,
	page_size int,
) utils.Blocks {
	response := http.Get(
		utils.BaseUrl+"blocks/"+block_id+"/children?page_size="+
			strconv.Itoa(page_size),
		secret_token,
	)
	var child_blocks utils.Blocks
	json.Unmarshal(response, &child_blocks)
	return child_blocks
}

// private func
func parseParagraphToMarkdown(block utils.Block) string {
	var paragraph_text string
	if len(block.Paragraph.Text) != 0 {
		paragraph_text = "\n" + block.Paragraph.Text[0].PlainText
	} else {
		paragraph_text = ""
	}
	return paragraph_text
}

func parseHeading1ToMarkdown(block utils.Block) string {
	var heading1_text string
	if len(block.Heading_1.Text) > 0 {
		heading1_text = "# " + block.Heading_1.Text[0].PlainText
	} else {
		heading1_text = "\n"
	}
	return heading1_text
}

func parseHeading2ToMarkdown(block utils.Block) string {
	var heading2_text string
	if len(block.Heading_2.Text) > 0 {
		heading2_text = "## " + block.Heading_2.Text[0].PlainText
	} else {
		heading2_text = "\n"
	}
	return heading2_text
}

func parseHeading3ToMarkdown(block utils.Block) string {
	var heading3_text string
	if len(block.Heading_3.Text) > 0 {
		heading3_text = "### " + block.Heading_3.Text[0].PlainText
	} else {
		heading3_text = "\n"
	}
	return heading3_text
}

func parseToDoToMarkdown(block utils.Block) string {
	var todo_text string
	if block.ToDo.Checked {
		checked := "- [X] "
		if len(block.ToDo.Text) > 0 {
			todo_text = checked + block.ToDo.Text[0].PlainText
		} else {
			todo_text = checked
		}
	} else {
		checked := "- [ ] "
		if len(block.ToDo.Text) > 0 {
			todo_text = checked + block.ToDo.Text[0].PlainText
		} else {
			todo_text = checked
		}
	}
	return todo_text
}

func parseChildPageToMarkdown(block utils.Block) string {
	var chiled_page_text string
	if len(block.ChildPage.Title) > 0 {
		chiled_page_text = ":page_facing_up: " + block.ChildPage.Title
	} else {
		chiled_page_text = ":page_facing_up: Untitled"
	}
	return chiled_page_text
}
