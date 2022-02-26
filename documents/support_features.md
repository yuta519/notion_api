# Supported Features
- DATABASES
  - [Retrieve a database](#Retrieve-a-database)
    - Return database information you specified.

  - [Retrieve a databases](#Retrieve-a-databases)
    - Return muliple databases information which are shared with your API.

  - [Retrieve a database IDs](#Retrieve-a-database-IDs)
    - Return database IDs in your Notion.

  - [Export Pages in Database](#Export-Pages-in-Database)
    - Export pages in database you specifiied to Markdown files.


- PAGES
  -

- BLOCKS
  -

- USERS
  -

- SEARCH
  -


# Usage Sample

## DATABASES

### Retrieve a database

- Sample Code
```golang
package main

import (
  "fmt"

	"github.com/yuta519/notion_api"
)

def main() {
	database := notion_api.FetchfDatabase(
		"secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion API key
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion Database ID
	)
	fmt.Println(databases)
}
```

### Retrieve a databases

- Sample code
```golang
package main

import (
  "fmt"

	"github.com/yuta519/notion_api"
)

def main() {
	database := notion_api.FetchfDatabases(
		"secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion API key
	)
	fmt.Println(databases)
}
```

- Sample outputs
```
[{<nil> 2022-01-26T22:41:00.000Z <nil> 1092f20d-09fb-407f-b36a-93b7a28c220a 2022-02-08T02:00:00.000Z database {workspace true} map[Assign:map[id:WbC%5D name:Assign people:map[] type:people] Name:map[id:title name:Name title:map[] type:title] Status:map[id:%3DrdM name:Status select:map[options:[map[color:red id:1 name:Not started] map[color:yellow id:2 name:In progress] map[color:green id:3 name:Completed]]] type:select]] [{map[bold:false code:false color:default italic:false strikethrough:false underline:false] <nil> ToDo map[content:ToDo link:] text}] https://www.notion.so/1092f20d09fb407fb36a93b7a28c220a 0xc00034c3c9} {<nil> 2022-01-25T23:15:00.000Z <nil> 0cc1608c-8be9-42ca-9bf8-3f6070d6c6d5 2022-01-26T22:32:00.000Z database {workspace true} map[Name:map[id:title name:Name title:map[] type:title] Tags:map[id:%3Abs%3B multi_select:map[options:[]] name:Tags type:multi_select]] [{map[bold:false code:false color:default italic:false strikethrough:false underline:false] <nil> aaa map[content:aaa link:] text}] https://www.notion.so/0cc1608c8be942ca9bf83f6070d6c6d5 0xc00034c49d} {<nil> 2022-01-30T23:51:00.000Z <nil> 17cb9b38-1749-46f0-9b86-8c2b77abd898 2022-02-18T00:31:00.000Z database {workspace true} map[Category:map[id:mAOk name:Category select:map[options:[]] type:select] Name:map[id:title name:Name title:map[] type:title] Status:map[id:bxnI name:Status select:map[options:[]] type:select]] [{map[bold:false code:false color:default italic:false strikethrough:false underline:false] <nil>  Blogs map[content: Blogs link:] text}] https://www.notion.so/17cb9b38174946f09b868c2b77abd898 0xc00034c59d}]
```


### Retrieve a database IDs
-  Sample code
```golang
package main

import (
  "fmt"

	"github.com/yuta519/notion_api"
)

def main() {
	databases := notion_api.FetchDatabaseIds(
		"secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion API key
	)
	fmt.Println(databases)
}
```

- Sample outputs
```
[map[id:1092f20d-09fb-407f-b36a-93b7a28c220a title:ToDo] map[id:0cc1608c-8be9-42ca-9bf8-3f6070d6c6d5 title:aaa] map[id:17cb9b38-1749-46f0-9b86-8c2b77abd898 title: Blogs]]
```

### Export Pages in Database

-  Sample code
```go
package main

import (
  "fmt"

	"github.com/yuta519/notion_api"
)

def main() {
	notion_api.ExportDbToMarkdown(
		"secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion API key
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // Please replace your Notion Database ID
	)
}
```

After executing this function, you will find created directory and md files.
