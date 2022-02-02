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

func parseTableToMarkdown(block utils.Block) string {
	var table_text string
	if block.Table.HasColumnHeader {
		fmt.Println()
	} else {
		table_text = "| | | \n| ---- | ---- | \n"
	}
	if block.Table.TableWidth > 0 {
		for i := 0; i < block.Table.TableWidth+1; i++ {
			table_text += "| "
		}
	} else {
		table_text = "\n"
	}
	return table_text
}

func parseBulletedListItemToMarkdown(block utils.Block) string {
	var bulleted_list_item string
	if len(block.BulletedListItem.Text) > 0 {
		bulleted_list_item = "- " + block.BulletedListItem.Text[0].PlainText
	} else {
		bulleted_list_item = "- "
	}
	return bulleted_list_item
}

func parseNumberedListItemToMarkdown(block utils.Block, num int) string {
	var numbered_list_item string
	if len(block.NumberedListItem.Text) > 0 {
		// for i := 1; i < len(block.NumberedListItem.Text)+1; i++ {
		numbered_list_item = strconv.Itoa(num) + ". " +
			block.NumberedListItem.Text[0].PlainText
		// }
	} else {
		numbered_list_item = "\n"
	}
	return numbered_list_item
}

func parseToggleToMarkdown(block utils.Block) string {
	var toggle string
	if len(block.Toggle.Text) > 0 {
		not_yet := "nothing"
		toggle = "<details><summary>" + block.Toggle.Text[0].PlainText +
			"</summary>" + not_yet + "</details>"
	} else {
		toggle = "\n"
	}
	return toggle
}

func parseQuoteToMarkdown(block utils.Block) string {
	var quote string
	if len(block.Quote.Text) > 0 {
		quote = "\n> " + block.Quote.Text[0].PlainText
	} else {
		quote = "\n"
	}
	return quote
}

func parseDividerToMarkdown(block utils.Block) string {
	return "---"
}

func parseCalloutToMarkdown(block utils.Block) string {
	callout := "```\n"
	if len(block.Callout.Text) > 0 {
		callout += block.Callout.Icon.Emoji + block.Callout.Text[0].PlainText +
			"\n```"
	} else {
		callout = "\n"
	}
	return callout

}
