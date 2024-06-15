package model

type ReqGetComments struct {
	PageSlug  string `form:"page_slug" binding:"required"`
	PageUrl   string `form:"page_url"`
	PageTitle string `form:"page_title"`
	LastID    string `form:"last_id"`
	ReplyTo   string `form:"reply_to"`
}

type RespGetComments struct {
	Comments []*Comment `json:"comments"`
}

type ReqAddComment struct {
	PageSlug string `json:"page_slug" form:"page_slug" binding:"required"`
	Content  string `json:"content" form:"content" binding:"required"`
	ReplyTo  string `json:"reply_to" form:"reply_to"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
}
