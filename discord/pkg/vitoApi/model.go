package vitoApi

type ContentType int

const (
	_ ContentType = iota
	Text
	Picture
	Video
	Sticker
	GIF
)

type Message struct {
	Content string      `json:"content"`
	Type    ContentType `json:"type"`
}

type Channel struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Sender struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
