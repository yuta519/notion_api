package utils

var (
	BaseUrl string = "https://api.notion.com/v1/"
)

type Objects struct {
	HasMore    bool        `json:"has_more"`
	NextCursor interface{} `json:"next_cursor"`
	Object     string      `json:"object"`
	Results    []Object    `json:"results"`
}

type Object struct {
	Cover          *Cover                 `json:"cover"`
	CreatedTime    string                 `json:"created_time"`
	Icon           *Icon                  `json:"icon"`
	Id             string                 `json:"id"`
	LastEditedTime string                 `json:"last_edited_time"`
	ObjectType     string                 `json:"object"`
	Parent         Parent                 `json:"parent"`
	Properties     map[string]interface{} `json:"properties"`
	Title          []Title                `json:"title"`
	Url            string                 `json:"url"`
}

type Cover struct {
	External map[string]string `json:"external"`
	Type     string            `json:"type"`
}

type Icon struct {
	Emoji string `json:"emoji"`
	Type  string `json:"type"`
}

type Parent struct {
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
}

type Title struct {
	Annotations map[string]interface{} `json:"annotations"`
	Href        *string                `json:"href"`
	PlainText   string                 `json:"plain_text"`
	Text        map[string]string      `json:"text"`
	Type        string                 `json:"type"`
}
