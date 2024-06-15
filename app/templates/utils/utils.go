package utils

import (
	"context"
	"encoding/json"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/supersupersimple/comment/app/ctxutil"
)

func ToHxVals(data map[string]any) string {
	res, _ := json.Marshal(data)
	return string(res)
}

func LoggedIn(ctx context.Context) bool {
	return ctxutil.LoggedIn(ctx)
}

func ToHtml(content string) string {
	unsafe := blackfriday.Run([]byte(content), blackfriday.WithExtensions(
		blackfriday.CommonExtensions|blackfriday.HardLineBreak,
	))
	sanitized := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return string(sanitized)
}
