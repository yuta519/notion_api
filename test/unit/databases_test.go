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

	res := notion_api.FetchRawResponseOfDatabases("")

	// To test a value of object key
	if res.ObjectType != "list" {
		t.Errorf("oh my godness")
	}
}
