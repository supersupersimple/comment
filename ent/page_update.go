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
)

// PageUpdate is the builder for updating Page entities.
type PageUpdate struct {
	config
	hooks    []Hook
	mutation *PageMutation
}

// Where appends a list predicates to the PageUpdate builder.
func (pu *PageUpdate) Where(ps ...predicate.Page) *PageUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetSlug sets the "slug" field.
func (pu *PageUpdate) SetSlug(s string) *PageUpdate {
	pu.mutation.SetSlug(s)
	return pu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (pu *PageUpdate) SetNillableSlug(s *string) *PageUpdate {
	if s != nil {
		pu.SetSlug(*s)
	}
	return pu
}

// SetTitle sets the "title" field.
func (pu *PageUpdate) SetTitle(s string) *PageUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pu *PageUpdate) SetNillableTitle(s *string) *PageUpdate {
	if s != nil {
		pu.SetTitle(*s)
	}
	return pu
}

// ClearTitle clears the value of the "title" field.
func (pu *PageUpdate) ClearTitle() *PageUpdate {
	pu.mutation.ClearTitle()
	return pu
}

// SetURL sets the "url" field.
func (pu *PageUpdate) SetURL(s string) *PageUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pu *PageUpdate) SetNillableURL(s *string) *PageUpdate {
	if s != nil {
		pu.SetURL(*s)
	}
	return pu
}

// ClearURL clears the value of the "url" field.
func (pu *PageUpdate) ClearURL() *PageUpdate {
	pu.mutation.ClearURL()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PageUpdate) SetCreatedAt(t time.Time) *PageUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PageUpdate) SetNillableCreatedAt(t *time.Time) *PageUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PageUpdate) AddCommentIDs(ids ...xid.ID) *PageUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PageUpdate) AddComments(c ...*Comment) *PageUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pu *PageUpdate) Mutation() *PageMutation {
	return pu.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PageUpdate) ClearComments() *PageUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PageUpdate) RemoveCommentIDs(ids ...xid.ID) *PageUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PageUpdate) RemoveComments(c ...*Comment) *PageUpdate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PageUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PageUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PageUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Slug(); ok {
		_spec.SetField(page.FieldSlug, field.TypeString, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if pu.mutation.TitleCleared() {
		_spec.ClearField(page.FieldTitle, field.TypeString)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(page.FieldURL, field.TypeString, value)
	}
	if pu.mutation.URLCleared() {
		_spec.ClearField(page.FieldURL, field.TypeString)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(page.FieldCreatedAt, field.TypeTime, value)
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
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
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PageUpdateOne is the builder for updating a single Page entity.
type PageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PageMutation
}

// SetSlug sets the "slug" field.
func (puo *PageUpdateOne) SetSlug(s string) *PageUpdateOne {
	puo.mutation.SetSlug(s)
	return puo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableSlug(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetSlug(*s)
	}
	return puo
}

// SetTitle sets the "title" field.
func (puo *PageUpdateOne) SetTitle(s string) *PageUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableTitle(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetTitle(*s)
	}
	return puo
}

// ClearTitle clears the value of the "title" field.
func (puo *PageUpdateOne) ClearTitle() *PageUpdateOne {
	puo.mutation.ClearTitle()
	return puo
}

// SetURL sets the "url" field.
func (puo *PageUpdateOne) SetURL(s string) *PageUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableURL(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetURL(*s)
	}
	return puo
}

// ClearURL clears the value of the "url" field.
func (puo *PageUpdateOne) ClearURL() *PageUpdateOne {
	puo.mutation.ClearURL()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PageUpdateOne) SetCreatedAt(t time.Time) *PageUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableCreatedAt(t *time.Time) *PageUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PageUpdateOne) AddCommentIDs(ids ...xid.ID) *PageUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PageUpdateOne) AddComments(c ...*Comment) *PageUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (puo *PageUpdateOne) Mutation() *PageMutation {
	return puo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PageUpdateOne) ClearComments() *PageUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PageUpdateOne) RemoveCommentIDs(ids ...xid.ID) *PageUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PageUpdateOne) RemoveComments(c ...*Comment) *PageUpdateOne {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the PageUpdate builder.
func (puo *PageUpdateOne) Where(ps ...predicate.Page) *PageUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PageUpdateOne) Select(field string, fields ...string) *PageUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Page entity.
func (puo *PageUpdateOne) Save(ctx context.Context) (*Page, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PageUpdateOne) SaveX(ctx context.Context) *Page {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PageUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PageUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PageUpdateOne) sqlSave(ctx context.Context) (_node *Page, err error) {
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Page.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for _, f := range fields {
			if !page.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != page.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Slug(); ok {
		_spec.SetField(page.FieldSlug, field.TypeString, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if puo.mutation.TitleCleared() {
		_spec.ClearField(page.FieldTitle, field.TypeString)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(page.FieldURL, field.TypeString, value)
	}
	if puo.mutation.URLCleared() {
		_spec.ClearField(page.FieldURL, field.TypeString)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(page.FieldCreatedAt, field.TypeTime, value)
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
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
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.CommentsTable,
			Columns: []string{page.CommentsColumn},
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
	_node = &Page{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
