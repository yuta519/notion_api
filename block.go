package notion_api

import (
	"encoding/json"
	"fmt"
	"os"
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
	err := json.Unmarshal(response, &child_blocks)
	if err != nil {
		os.Exit(1)
	}
	return child_blocks
}
