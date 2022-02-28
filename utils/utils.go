package utils

var (
	BaseUrl string = "https://api.notion.com/v1/"
)

type Objects struct {
	HasMore    bool        `json:"has_more"`
	NextCursor interface{} `json:"next_cursor"`
	ObjectType string      `json:"object"`
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
	Archived       *bool                  `json:"archived"`
}

type Blocks struct {
	HasMore    bool        `json:"has_more"`
	NextCursor interface{} `json:"next_cursor"`
	ObjectType string      `json:"object"`
	Results    []Block     `json:"results"`
}

type Block struct {
	ObjectType       string            `json:"object"`
	Id               string            `json:"id"`
	CreatedTime      string            `json:"created_time"`
	LastEditedTime   string            `json:"last_edited_time"`
	HasChildren      bool              `json:"has_children"`
	Archived         bool              `json:"archived"`
	Type             string            `json:"type"`
	Paragraph        *Paragraph        `json:"paragraph"`
	Heading_1        *Heading          `json:"heading_1"`
	Heading_2        *Heading          `json:"heading_2"`
	Heading_3        *Heading          `json:"heading_3"`
	ToDo             *ToDo             `json:"to_do"`
	ChildPage        *ChildPage        `json:"child_page"`
	Table            *Table            `json:"table"`
	TableRow         *TableRow         `json:"table_row"`
	BulletedListItem *BulletedListItem `json:"bulleted_list_item"`
	NumberedListItem *NumberedListItem `json:"numbered_list_item"`
	Toggle           *Toggle           `json:"toggle"`
	Quote            *Quote            `json:"quote"`
	Divider          *interface{}      `json:"divider"`
	Callout          *Callout          `json:"callout"`
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

type TextBlock struct {
	Annotations map[string]interface{} `json:"annotations"`
	Href        *string                `json:"href"`
	PlainText   string                 `json:"plain_text"`
	Text        map[string]string      `json:"text"`
	Type        string                 `json:"type"`
}

type Paragraph struct {
	Text []TextBlock `json:"text"`
}

type Heading struct {
	Text []TextBlock `json:"text"`
}

type ToDo struct {
	Text    []TextBlock `json:"text"`
	Checked bool        `json:"checked"`
}

type ChildPage struct {
	Title string `json:"title"`
}

type Table struct {
	TableWidth      int  `json:"table_width"`
	HasColumnHeader bool `json:"has_column_header"`
	HasRowHeader    bool `json:"has_row_header"`
}

type TableRow struct {
	Cells [][]Cell `json:"cells"`
}

type Cell struct {
	Type        string      `json:"type"`
	Text        interface{} `json:"text"`
	Annotations interface{} `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type BulletedListItem struct {
	Text []TextBlock `json:"text"`
}

type NumberedListItem struct {
	Text []TextBlock `json:"text"`
}

type Toggle struct {
	Text []TextBlock `json:"text"`
}

type Quote struct {
	Text []TextBlock `json:"text"`
}

type Callout struct {
	Text []TextBlock `json:"text"`
	Icon Icon        `json:"icon"`
}
