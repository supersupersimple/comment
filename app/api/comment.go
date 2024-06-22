package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"github.com/rs/xid"
	"github.com/supersupersimple/comment/app/ctxutil"
	"github.com/supersupersimple/comment/app/model"
	"github.com/supersupersimple/comment/app/render"
	"github.com/supersupersimple/comment/app/templates/comment/components"
	"github.com/supersupersimple/comment/app/templates/comment/pages"
	"github.com/supersupersimple/comment/ent"
	"github.com/supersupersimple/comment/ent/comment"
	"github.com/supersupersimple/comment/ent/page"
	"github.com/supersupersimple/comment/ent/user"
)

func errorReplyBox(c *gin.Context, replyTo string, errorStatus model.ReplyStatus) {
	render.HTML(c, http.StatusOK, components.ReplyBox(&model.Comment{ID: replyTo}, errorStatus))
}

func AddComment(c *gin.Context) {
	var req = &model.ReqAddComment{}
	if err := c.ShouldBind(req); err != nil {
		errorReplyBox(c, req.ReplyTo, model.ReplyStatusInvalidReq)
		return
	}

	client := ctxutil.DB(c)
	pa, err := getPage(c, req.PageSlug, "", "")
	if err != nil {
		errorReplyBox(c, req.ReplyTo, model.ReplyStatusFailed)
		return
	}
	u, err := getUser(c, req.Username, req.Email)
	if err != nil {
		errorReplyBox(c, req.ReplyTo, model.ReplyStatusFailed)
		return
	}
	depth := 0
	var replyToID *xid.ID
	if req.ReplyTo != "" {
		re, err := getComment(c, req.ReplyTo)
		if err != nil {
			errorReplyBox(c, req.ReplyTo, model.ReplyStatusFailed)
			return
		}
		replyToID = &re.ID
		depth = re.Depth + 1
	}

	status := comment.StatusPending
	if u.IsAdmin {
		status = comment.StatusApproved
	}
	_, err = client.Comment.Create().
		SetContent(req.Content).
		SetPage(pa).
		SetUser(u).
		SetDepth(depth).
		SetNillableReplyToID(replyToID).
		SetStatus(status).
		Save(c)

	if err != nil {
		errorReplyBox(c, req.ReplyTo, model.ReplyStatusFailed)
		return
	}

	render.HTML(c, http.StatusCreated, components.ReplyBox(&model.Comment{ID: req.ReplyTo}, model.ReplyStatusSuccessful))
}

func GetComments(c *gin.Context) {
	req := &model.ReqGetComments{}
	if err := c.ShouldBind(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pa, err := getPage(c, req.PageSlug, req.PageUrl, req.PageTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	var replyTo xid.ID
	if req.ReplyTo == "" {
		replyTo = xid.NilID()
	} else {
		replyTo, err = xid.FromString(req.ReplyTo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
	comments, err := getCommentsTree(c, ctxutil.DB(c), pa.ID, replyTo, req.LastID, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if req.LastID == "" && req.ReplyTo == "" {
		render.HTML(c, http.StatusOK, pages.Page(&model.PageVals{
			PageSlug:  req.PageSlug,
			CsrfToken: nosurf.Token(c.Request),
		}, comments))
	} else {
		render.HTML(c, http.StatusOK, components.CommentList(comments))
	}
}

func getCommentsTree(c *gin.Context, client *ent.Client, pageID int64, replyToID xid.ID, lastID string, loopDepth int) ([]*model.Comment, error) {
	if loopDepth > model.MaxLoopDepth {
		return nil, nil
	}
	cq := client.Comment.Query().Where(
		comment.PageIDEQ(pageID),
	)
	if ctxutil.LoggedIn(c) {
		cq.Where(comment.StatusIn(comment.StatusPending, comment.StatusApproved))
	} else {
		cq.Where(comment.StatusEQ(comment.StatusApproved))
	}
	if lastID != "" {
		lid, err := xid.FromString(lastID)
		if err != nil {
			return nil, err
		}
		cq.Where(comment.IDGT(lid))
	}
	if replyToID.IsNil() {
		cq.Where(comment.ReplyToIDIsNil())
	} else {
		cq.Where(comment.ReplyToIDEQ(replyToID))
	}
	replies, err := cq.WithUser().
		Order(ent.Asc(comment.FieldID)).
		Limit(model.LimitPerBatch).
		All(c)
	if err != nil {
		return nil, err
	}

	tree := make([]*model.Comment, len(replies))
	for i, reply := range replies {
		children, err := getCommentsTree(c, client, pageID, reply.ID, "", loopDepth+1)
		if err != nil {
			return nil, err
		}
		tree[i] = &model.Comment{
			ID:          reply.ID.String(),
			Username:    reply.Edges.User.Username,
			Content:     reply.Content,
			PublishTime: ConvertPublishedTime(reply.CreatedAt),
			Depth:       reply.Depth,
			Children:    children,
			LastDepth:   loopDepth == model.MaxLoopDepth,
			ReplyTo:     replyToID.String(),
			IsAdmin:     reply.Edges.User.IsAdmin,
			Status:      reply.Status.String(),
		}
	}
	return tree, nil
}

// get page if not exist, create
func getPage(c *gin.Context, slug, url, title string) (*ent.Page, error) {
	client := ctxutil.DB(c)
	pa, err := client.Page.Query().Where(
		page.SlugEQ(slug),
	).Only(c)
	if ent.IsNotFound(err) {
		// create page
		pa, err = client.Page.Create().SetSlug(slug).SetURL(url).SetTitle(title).Save(c)
		if err != nil {
			return nil, err
		}
	}
	return pa, err
}

// get user if not exist, create
func getUser(c *gin.Context, username, email string) (u *ent.User, err error) {
	client := ctxutil.DB(c)
	if ctxutil.LoggedIn(c) {
		u, err = client.User.Query().Where(
			user.IsAdminEQ(true),
		).Only(c)
	} else {
		u, err = client.User.Query().Where(
			user.UsernameEQ(username),
			user.EmailEQ(email),
			user.IsAdminEQ(false),
		).Only(c)
		if ent.IsNotFound(err) {
			// create user
			u, err = client.User.Create().SetUsername(username).SetEmail(email).Save(c)
			if err != nil {
				return nil, err
			}
		}
	}
	return u, err
}

func getComment(c *gin.Context, id string) (*ent.Comment, error) {
	cid, err := xid.FromString(id)
	if err != nil {
		return nil, err
	}
	client := ctxutil.DB(c)
	return client.Comment.Query().Where(
		comment.IDEQ(cid),
	).Only(c)
}

func toCommentModel(co *ent.Comment) *model.Comment {
	replyTo := ""
	if co.Edges.ReplyTo != nil {
		replyTo = co.Edges.ReplyTo.ID.String()
	}
	return &model.Comment{
		ID:          co.ID.String(),
		Username:    co.Edges.User.Username,
		Content:     co.Content,
		PublishTime: ConvertPublishedTime(co.CreatedAt),
		Depth:       co.Depth,
		Children:    nil,
		LastDepth:   false,
		ReplyTo:     replyTo,
		IsAdmin:     co.Edges.User.IsAdmin,
		Status:      co.Status.String(),
	}
}
