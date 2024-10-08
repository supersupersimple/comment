// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
	"github.com/supersupersimple/comment/ent/comment"
	"github.com/supersupersimple/comment/ent/page"
	"github.com/supersupersimple/comment/ent/predicate"
	"github.com/supersupersimple/comment/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// SetStatus sets the "status" field.
func (cu *CommentUpdate) SetStatus(c comment.Status) *CommentUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableStatus(c *comment.Status) *CommentUpdate {
	if c != nil {
		cu.SetStatus(*c)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDepth sets the "depth" field.
func (cu *CommentUpdate) SetDepth(i int) *CommentUpdate {
	cu.mutation.ResetDepth()
	cu.mutation.SetDepth(i)
	return cu
}

// SetNillableDepth sets the "depth" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDepth(i *int) *CommentUpdate {
	if i != nil {
		cu.SetDepth(*i)
	}
	return cu
}

// AddDepth adds i to the "depth" field.
func (cu *CommentUpdate) AddDepth(i int) *CommentUpdate {
	cu.mutation.AddDepth(i)
	return cu
}

// SetApproveToken sets the "approve_token" field.
func (cu *CommentUpdate) SetApproveToken(s string) *CommentUpdate {
	cu.mutation.SetApproveToken(s)
	return cu
}

// SetNillableApproveToken sets the "approve_token" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableApproveToken(s *string) *CommentUpdate {
	if s != nil {
		cu.SetApproveToken(*s)
	}
	return cu
}

// ClearApproveToken clears the value of the "approve_token" field.
func (cu *CommentUpdate) ClearApproveToken() *CommentUpdate {
	cu.mutation.ClearApproveToken()
	return cu
}

// SetPageID sets the "page_id" field.
func (cu *CommentUpdate) SetPageID(i int64) *CommentUpdate {
	cu.mutation.SetPageID(i)
	return cu
}

// SetNillablePageID sets the "page_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillablePageID(i *int64) *CommentUpdate {
	if i != nil {
		cu.SetPageID(*i)
	}
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommentUpdate) SetUserID(i int64) *CommentUpdate {
	cu.mutation.SetUserID(i)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUserID(i *int64) *CommentUpdate {
	if i != nil {
		cu.SetUserID(*i)
	}
	return cu
}

// SetReplyToID sets the "reply_to_id" field.
func (cu *CommentUpdate) SetReplyToID(x xid.ID) *CommentUpdate {
	cu.mutation.SetReplyToID(x)
	return cu
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableReplyToID(x *xid.ID) *CommentUpdate {
	if x != nil {
		cu.SetReplyToID(*x)
	}
	return cu
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cu *CommentUpdate) ClearReplyToID() *CommentUpdate {
	cu.mutation.ClearReplyToID()
	return cu
}

// SetPage sets the "page" edge to the Page entity.
func (cu *CommentUpdate) SetPage(p *Page) *CommentUpdate {
	return cu.SetPageID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cu *CommentUpdate) SetUser(u *User) *CommentUpdate {
	return cu.SetUserID(u.ID)
}

// SetReplyTo sets the "reply_to" edge to the Comment entity.
func (cu *CommentUpdate) SetReplyTo(c *Comment) *CommentUpdate {
	return cu.SetReplyToID(c.ID)
}

// AddReplyIDs adds the "replies" edge to the Comment entity by IDs.
func (cu *CommentUpdate) AddReplyIDs(ids ...xid.ID) *CommentUpdate {
	cu.mutation.AddReplyIDs(ids...)
	return cu
}

// AddReplies adds the "replies" edges to the Comment entity.
func (cu *CommentUpdate) AddReplies(c ...*Comment) *CommentUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddReplyIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearPage clears the "page" edge to the Page entity.
func (cu *CommentUpdate) ClearPage() *CommentUpdate {
	cu.mutation.ClearPage()
	return cu
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CommentUpdate) ClearUser() *CommentUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearReplyTo clears the "reply_to" edge to the Comment entity.
func (cu *CommentUpdate) ClearReplyTo() *CommentUpdate {
	cu.mutation.ClearReplyTo()
	return cu
}

// ClearReplies clears all "replies" edges to the Comment entity.
func (cu *CommentUpdate) ClearReplies() *CommentUpdate {
	cu.mutation.ClearReplies()
	return cu
}

// RemoveReplyIDs removes the "replies" edge to Comment entities by IDs.
func (cu *CommentUpdate) RemoveReplyIDs(ids ...xid.ID) *CommentUpdate {
	cu.mutation.RemoveReplyIDs(ids...)
	return cu
}

// RemoveReplies removes "replies" edges to Comment entities.
func (cu *CommentUpdate) RemoveReplies(c ...*Comment) *CommentUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveReplyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if _, ok := cu.mutation.PageID(); cu.mutation.PageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Comment.page"`)
	}
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Comment.user"`)
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Depth(); ok {
		_spec.SetField(comment.FieldDepth, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedDepth(); ok {
		_spec.AddField(comment.FieldDepth, field.TypeInt, value)
	}
	if value, ok := cu.mutation.ApproveToken(); ok {
		_spec.SetField(comment.FieldApproveToken, field.TypeString, value)
	}
	if cu.mutation.ApproveTokenCleared() {
		_spec.ClearField(comment.FieldApproveToken, field.TypeString)
	}
	if cu.mutation.PageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PageTable,
			Columns: []string{comment.PageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PageTable,
			Columns: []string{comment.PageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ReplyToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ReplyToTable,
			Columns: []string{comment.ReplyToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ReplyToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ReplyToTable,
			Columns: []string{comment.ReplyToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !cu.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CommentUpdateOne) SetStatus(c comment.Status) *CommentUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableStatus(c *comment.Status) *CommentUpdateOne {
	if c != nil {
		cuo.SetStatus(*c)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDepth sets the "depth" field.
func (cuo *CommentUpdateOne) SetDepth(i int) *CommentUpdateOne {
	cuo.mutation.ResetDepth()
	cuo.mutation.SetDepth(i)
	return cuo
}

// SetNillableDepth sets the "depth" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDepth(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetDepth(*i)
	}
	return cuo
}

// AddDepth adds i to the "depth" field.
func (cuo *CommentUpdateOne) AddDepth(i int) *CommentUpdateOne {
	cuo.mutation.AddDepth(i)
	return cuo
}

// SetApproveToken sets the "approve_token" field.
func (cuo *CommentUpdateOne) SetApproveToken(s string) *CommentUpdateOne {
	cuo.mutation.SetApproveToken(s)
	return cuo
}

// SetNillableApproveToken sets the "approve_token" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableApproveToken(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetApproveToken(*s)
	}
	return cuo
}

// ClearApproveToken clears the value of the "approve_token" field.
func (cuo *CommentUpdateOne) ClearApproveToken() *CommentUpdateOne {
	cuo.mutation.ClearApproveToken()
	return cuo
}

// SetPageID sets the "page_id" field.
func (cuo *CommentUpdateOne) SetPageID(i int64) *CommentUpdateOne {
	cuo.mutation.SetPageID(i)
	return cuo
}

// SetNillablePageID sets the "page_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillablePageID(i *int64) *CommentUpdateOne {
	if i != nil {
		cuo.SetPageID(*i)
	}
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CommentUpdateOne) SetUserID(i int64) *CommentUpdateOne {
	cuo.mutation.SetUserID(i)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUserID(i *int64) *CommentUpdateOne {
	if i != nil {
		cuo.SetUserID(*i)
	}
	return cuo
}

// SetReplyToID sets the "reply_to_id" field.
func (cuo *CommentUpdateOne) SetReplyToID(x xid.ID) *CommentUpdateOne {
	cuo.mutation.SetReplyToID(x)
	return cuo
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableReplyToID(x *xid.ID) *CommentUpdateOne {
	if x != nil {
		cuo.SetReplyToID(*x)
	}
	return cuo
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cuo *CommentUpdateOne) ClearReplyToID() *CommentUpdateOne {
	cuo.mutation.ClearReplyToID()
	return cuo
}

// SetPage sets the "page" edge to the Page entity.
func (cuo *CommentUpdateOne) SetPage(p *Page) *CommentUpdateOne {
	return cuo.SetPageID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CommentUpdateOne) SetUser(u *User) *CommentUpdateOne {
	return cuo.SetUserID(u.ID)
}

// SetReplyTo sets the "reply_to" edge to the Comment entity.
func (cuo *CommentUpdateOne) SetReplyTo(c *Comment) *CommentUpdateOne {
	return cuo.SetReplyToID(c.ID)
}

// AddReplyIDs adds the "replies" edge to the Comment entity by IDs.
func (cuo *CommentUpdateOne) AddReplyIDs(ids ...xid.ID) *CommentUpdateOne {
	cuo.mutation.AddReplyIDs(ids...)
	return cuo
}

// AddReplies adds the "replies" edges to the Comment entity.
func (cuo *CommentUpdateOne) AddReplies(c ...*Comment) *CommentUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddReplyIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearPage clears the "page" edge to the Page entity.
func (cuo *CommentUpdateOne) ClearPage() *CommentUpdateOne {
	cuo.mutation.ClearPage()
	return cuo
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CommentUpdateOne) ClearUser() *CommentUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearReplyTo clears the "reply_to" edge to the Comment entity.
func (cuo *CommentUpdateOne) ClearReplyTo() *CommentUpdateOne {
	cuo.mutation.ClearReplyTo()
	return cuo
}

// ClearReplies clears all "replies" edges to the Comment entity.
func (cuo *CommentUpdateOne) ClearReplies() *CommentUpdateOne {
	cuo.mutation.ClearReplies()
	return cuo
}

// RemoveReplyIDs removes the "replies" edge to Comment entities by IDs.
func (cuo *CommentUpdateOne) RemoveReplyIDs(ids ...xid.ID) *CommentUpdateOne {
	cuo.mutation.RemoveReplyIDs(ids...)
	return cuo
}

// RemoveReplies removes "replies" edges to Comment entities.
func (cuo *CommentUpdateOne) RemoveReplies(c ...*Comment) *CommentUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveReplyIDs(ids...)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.PageID(); cuo.mutation.PageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Comment.page"`)
	}
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Comment.user"`)
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Depth(); ok {
		_spec.SetField(comment.FieldDepth, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedDepth(); ok {
		_spec.AddField(comment.FieldDepth, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.ApproveToken(); ok {
		_spec.SetField(comment.FieldApproveToken, field.TypeString, value)
	}
	if cuo.mutation.ApproveTokenCleared() {
		_spec.ClearField(comment.FieldApproveToken, field.TypeString)
	}
	if cuo.mutation.PageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PageTable,
			Columns: []string{comment.PageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PageTable,
			Columns: []string{comment.PageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ReplyToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ReplyToTable,
			Columns: []string{comment.ReplyToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ReplyToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ReplyToTable,
			Columns: []string{comment.ReplyToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRepliesIDs(); len(nodes) > 0 && !cuo.mutation.RepliesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
