package notion_api

import (
	"fmt"
	"strconv"

	"github.com/yuta519/notion_api/handler/http"
	"github.com/yuta519/notion_api/utils"
)

func FetchBlock(secret_token string, block_id string) {
	response := http.Get(utils.BaseUrl+"blocks/"+block_id, secret_token)
	fmt.Println(string(response))
}

func FetchChildrenBlock(secret_token string, block_id string, page_size int) {
	response := http.Get(
		utils.BaseUrl+"blocks/"+block_id+"/children?page_size="+
			strconv.Itoa(page_size),
		secret_token,
	)
	fmt.Println(string(response))
}
