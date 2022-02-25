# Supported Features
- DATABASES
  - [Retrieve a database](#Retrieve-a-database)
    - Return database information you specified.

  - [Retrieve a database IDs](#Retrieve-a-database-IDs)
    - Return database IDs in your Notion.



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

### Retrieve a database IDs
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
