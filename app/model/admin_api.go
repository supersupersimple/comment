package model

type AdminPageVals struct {
	CsrfToken string
	Source    string
	ErrorMsg  string
}

const (
	AdminLimitPerBatch = 20
	SourceAdmin        = "admin"
)

type StatusFilter uint

const (
	StatusPending StatusFilter = iota
	StatusAll
)

type ReqAdminGetComments struct {
	LastID string `form:"last_id"`

	StatusFilter uint `form:"status_filter"`
}

type ReqAdminLogin struct {
	Password string `form:"password" binding:"required"`
}

type ReqAdminSetup struct {
	Name          string `form:"name" binding:"required,min=1"`     // name length must be longer than 1
	Password      string `form:"password" binding:"required,min=8"` // password length must be longer than 8
	AllowOrigins  string `form:"allow_origins" binding:"required"`
	LimitPerBatch int    `form:"limit_per_batch" binding:"required,gte=5,lte=20"`
	MaxLoopDepth  int    `form:"max_loop_depth" binding:"required,gte=3,lte=5"`
}

type AdminComment struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Content     string `json:"content"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PageSlug    string `json:"page_slug"`
	PageTitle   string `json:"page_title"`
	PublishTime string `json:"publish_time"`
	Depth       int    `json:"depth"`
	ReplyTo     string `json:"reply_to"`
}
