package api

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"github.com/supersupersimple/comment/app/config"
	"github.com/supersupersimple/comment/app/ctxutil"
	"github.com/supersupersimple/comment/app/model"
	"github.com/supersupersimple/comment/app/render"
	"github.com/supersupersimple/comment/app/templates/admin/admincomponents"
	"github.com/supersupersimple/comment/app/templates/admin/adminpages"
	"github.com/supersupersimple/comment/app/templates/comment/components"
	"github.com/supersupersimple/comment/ent"
	"github.com/supersupersimple/comment/ent/comment"
)

func AdminGetComments(c *gin.Context) {
	req := &model.ReqAdminGetComments{}
	if err := c.ShouldBind(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	client := ctxutil.DB(c)
	cq := client.Comment.Query()
	if req.LastID != "" {
		lastID, err := xid.FromString(req.LastID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		cq.Where(comment.IDGT(lastID))
	}
	switch req.StatusFilter {
	case uint(model.StatusPending):
		cq.Where(comment.StatusEQ(comment.StatusPending))
	}
	comments, err := cq.WithPage().WithUser().
		Order(ent.Asc(comment.FieldID)).Limit(model.AdminLimitPerBatch).All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var adminComments []*model.AdminComment
	for _, c := range comments {
		ac := toAdminComment(c)
		adminComments = append(adminComments, ac)
	}
	render.HTML(c, http.StatusOK, adminpages.Page(&model.AdminPageVals{
		CsrfToken: nosurf.Token(c.Request),
		Source:    model.SourceAdmin,
	}, adminComments,
	))
}

func ApproveComment(c *gin.Context) {
	id, err := xid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	client := ctxutil.DB(c)
	err = client.Comment.UpdateOneID(id).SetStatus(comment.StatusApproved).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	comment, err := client.Comment.Query().Where(comment.IDEQ(id)).WithPage().WithUser().WithReplyTo().Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if c.PostForm("source") == model.SourceAdmin {
		render.HTML(c, http.StatusAccepted, admincomponents.CommentRow(toAdminComment(comment)))
	} else {
		render.HTML(c, http.StatusOK, components.CommentBox(toCommentModel(comment)))
	}
}

func RejectComment(c *gin.Context) {
	id, err := xid.FromString(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	client := ctxutil.DB(c)
	err = client.Comment.UpdateOneID(id).SetStatus(comment.StatusRejected).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	comment, err := client.Comment.Query().Where(comment.IDEQ(id)).WithPage().WithUser().Only(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if c.Query("source") == model.SourceAdmin {
		render.HTML(c, http.StatusAccepted, admincomponents.CommentRow(toAdminComment(comment)))
	}
}

func AdminSetup(c *gin.Context) {
	if viper.GetBool(config.Setup) {
		return
	}
	if c.Request.Method == http.MethodGet {
		render.HTML(c, http.StatusOK, adminpages.Setup(&model.AdminPageVals{
			CsrfToken: nosurf.Token(c.Request),
		}))
		return
	}
	if c.Request.Method != http.MethodPost {
		return
	}
	// POST
	req := &model.ReqAdminSetup{}
	if err := c.ShouldBind(req); err != nil {
		render.HTML(c, http.StatusBadRequest, adminpages.Setup(&model.AdminPageVals{
			CsrfToken: nosurf.Token(c.Request),
			ErrorMsg:  err.Error(),
		}))
		return
	}
	hashPwd, err := HashPassword(req.Password)
	if err != nil {
		render.HTML(c, http.StatusBadRequest, adminpages.Setup(&model.AdminPageVals{
			CsrfToken: nosurf.Token(c.Request),
			ErrorMsg:  "invalid password",
		}))
		return
	}
	origins := strings.Split(req.AllowOrigins, "\n")

	err = WithTx(c, ctxutil.DB(c), func(tx *ent.Tx) error {
		cfg := tx.Conf.Create()
		cfg.SetCookieSecret(viper.GetString(config.Session))
		cfg.SetLimitPerBatch(req.LimitPerBatch)
		cfg.SetMaxLoopDepth(req.MaxLoopDepth)
		cfg.SetPassword(hashPwd)
		cfg.SetAllowOrigins(origins)
		_, err := cfg.Save(c)
		if err != nil {
			return err
		}
		u := tx.User.Create().SetIsAdmin(true)
		u.SetUsername(req.Name)
		_, err = u.Save(c)
		return err
	})
	if err != nil {
		render.HTML(c, http.StatusBadRequest, adminpages.Setup(&model.AdminPageVals{
			CsrfToken: nosurf.Token(c.Request),
			ErrorMsg:  "db error",
		}))
		return
	}
	config.LoadConfig(ctxutil.DB(c))
	c.Redirect(http.StatusFound, "/admin/login")
}

func AdminLogin(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		render.HTML(c, http.StatusOK, adminpages.Login(&model.AdminPageVals{
			CsrfToken: nosurf.Token(c.Request),
		}))
		return
	}
	if c.Request.Method == http.MethodPost {
		req := &model.ReqAdminLogin{}
		if err := c.ShouldBind(req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// authenticate admin credentials
		// assume you have a function to authenticate admin credentials
		conf := ctxutil.DB(c).Conf.Query().OnlyX(c)
		if !CheckPasswordHash(req.Password, conf.Password) {
			c.String(http.StatusUnauthorized, "Invalid password")
			return
		}

		// login successful, set session or cookie
		session := sessions.Default(c)
		session.Set("user", true)
		session.Save()

		c.Redirect(http.StatusFound, "/admin/comments")
		return
	}
}

func toAdminComment(c *ent.Comment) *model.AdminComment {
	replyTo := ""
	if c.Edges.ReplyTo != nil {
		replyTo = c.Edges.ReplyTo.ID.String()
	}
	ac := &model.AdminComment{
		ID:          c.ID.String(),
		Status:      c.Status.String(),
		Content:     c.Content,
		PublishTime: c.CreatedAt.Format("2006-01-02 15:04:05"),
		Username:    c.Edges.User.Username,
		Email:       c.Edges.User.Email,
		PageSlug:    c.Edges.Page.Slug,
		PageTitle:   c.Edges.Page.Title,
		Depth:       c.Depth,
		ReplyTo:     replyTo,
	}
	return ac
}
