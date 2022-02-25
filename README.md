# Notion API

## About
This is a notion api SDK of golang.

## WHY I created this library
Do you use [notion](https://notion.so)?
I love to use a Notion.

So I, heavy Notion user, want to use Notion anywhere.

But some reasons distrub this desire.

For example:
- Notion is prohibited in a company.
- Using notion makes it difficult to export knowledge when you need.

## Techniques
This library made of:
- Golang

## Easy to use
This section expains how to use this library.

1. Download the library.

```
 $ go get github.com/yuta519/notion_api
 $ go mod tidy
```

2. Create sample go file.
``` golang
$ touch main.go
```

3. Edit `main.go` like below.
```go
package main

import (
    "fmt"

	"github.com/yuta519/notion_api"
)

func main() {
    db_ids := notion_api.FetchDatabaseIds(
        "<notion API Secret Key>",
	)
	fmt.Println(db_ids)
  }
```


4. Run code!
```
go run main.go
```

If it works well, you'll get response like below.
(You can get an id of your database and title.)
```
map[id:17cb9b38-1749-46f0-9b86-8c2b77abd898 title:API Test]]
```


<!-- # Architecture -->

<!-- # Upcoming Features -->

## License
Copyright 2022 Yuta Kawamura

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
