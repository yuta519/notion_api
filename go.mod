module github.com/yuta519/notion_api

go 1.17

replace github.com/yuta519/notion_api/lib => ./lib

replace github.com/yuta519/notion_api/handler/http => ./handler/http/

replace github.com/yuta519/notion_api/utils => ./utils

require github.com/jarcoal/httpmock v1.1.0
