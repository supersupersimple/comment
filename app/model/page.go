package model

const (
	LimitPerBatch = 5 // TODO put to config
	MaxLoopDepth  = 3
)

type PageVals struct {
	PageSlug  string
	CsrfToken string
}

type ReplyStatus string

const (
	ReplyStatusSuccessful  ReplyStatus = "Your comment has been sent. Please wait for approval."
	ReplyStatusFailed      ReplyStatus = "Your comment sent failed. Please try again."
	ReplyStatusInvalidReq  ReplyStatus = "Invalid request. Please fill name and reply content."
	ReplyStatusInvalidCSRF ReplyStatus = "Invalid CSRF token. Please refresh page and try again."
)
