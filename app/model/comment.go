package model

type Comment struct {
	ID          string     `json:"id"`
	Content     string     `json:"content"`
	Username    string     `json:"username"`
	PublishTime string     `json:"publish_time"`
	Depth       int        `json:"depth"`
	Children    []*Comment `json:"children"`
	LastDepth   bool       `json:"last_depth"`
	ReplyTo     string     `json:"reply_to"`
	IsAdmin     bool       `json:"is_admin"`
	Status      string     `json:"status"`
}
