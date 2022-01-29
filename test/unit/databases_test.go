package unit

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/yuta519/notion_api"
	"github.com/yuta519/notion_api/utils"
)

func TestFetchDatabaseIds(t *testing.T) {
	// Mock Notion API response
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		utils.BaseUrl+"databases",
		httpmock.NewBytesResponder(
			200,
			[]byte("{\"object\":\"list\",\"results\":[],\"next_cursor\":null,\"has_more\":false}")),
	)

	res := notion_api.FetchDatabaseIds("")

	// To test correct parsing
	_, objectOk := res["object"]
	_, resultsOk := res["results"]
	_, nextcursorOk := res["next_cursor"]
	_, hasMoreOk := res["has_more"]
	if !objectOk || !resultsOk || !nextcursorOk || !hasMoreOk {
		t.Errorf("Some fileds are parsed incorrectly.")
	}

	// To test a value of object key
	if res["object"].(string) != "list" {
		t.Errorf("oh my godness")
	}
}
