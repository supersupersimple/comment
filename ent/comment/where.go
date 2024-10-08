// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
	"github.com/supersupersimple/comment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// Depth applies equality check predicate on the "depth" field. It's identical to DepthEQ.
func Depth(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldDepth, v))
}

// ApproveToken applies equality check predicate on the "approve_token" field. It's identical to ApproveTokenEQ.
func ApproveToken(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldApproveToken, v))
}

// PageID applies equality check predicate on the "page_id" field. It's identical to PageIDEQ.
func PageID(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldPageID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// ReplyToID applies equality check predicate on the "reply_to_id" field. It's identical to ReplyToIDEQ.
func ReplyToID(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldReplyToID, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldContent, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUpdatedAt, v))
}

// DepthEQ applies the EQ predicate on the "depth" field.
func DepthEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldDepth, v))
}

// DepthNEQ applies the NEQ predicate on the "depth" field.
func DepthNEQ(v int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldDepth, v))
}

// DepthIn applies the In predicate on the "depth" field.
func DepthIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldDepth, vs...))
}

// DepthNotIn applies the NotIn predicate on the "depth" field.
func DepthNotIn(vs ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldDepth, vs...))
}

// DepthGT applies the GT predicate on the "depth" field.
func DepthGT(v int) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldDepth, v))
}

// DepthGTE applies the GTE predicate on the "depth" field.
func DepthGTE(v int) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldDepth, v))
}

// DepthLT applies the LT predicate on the "depth" field.
func DepthLT(v int) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldDepth, v))
}

// DepthLTE applies the LTE predicate on the "depth" field.
func DepthLTE(v int) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldDepth, v))
}

// ApproveTokenEQ applies the EQ predicate on the "approve_token" field.
func ApproveTokenEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldApproveToken, v))
}

// ApproveTokenNEQ applies the NEQ predicate on the "approve_token" field.
func ApproveTokenNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldApproveToken, v))
}

// ApproveTokenIn applies the In predicate on the "approve_token" field.
func ApproveTokenIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldApproveToken, vs...))
}

// ApproveTokenNotIn applies the NotIn predicate on the "approve_token" field.
func ApproveTokenNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldApproveToken, vs...))
}

// ApproveTokenGT applies the GT predicate on the "approve_token" field.
func ApproveTokenGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldApproveToken, v))
}

// ApproveTokenGTE applies the GTE predicate on the "approve_token" field.
func ApproveTokenGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldApproveToken, v))
}

// ApproveTokenLT applies the LT predicate on the "approve_token" field.
func ApproveTokenLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldApproveToken, v))
}

// ApproveTokenLTE applies the LTE predicate on the "approve_token" field.
func ApproveTokenLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldApproveToken, v))
}

// ApproveTokenContains applies the Contains predicate on the "approve_token" field.
func ApproveTokenContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldApproveToken, v))
}

// ApproveTokenHasPrefix applies the HasPrefix predicate on the "approve_token" field.
func ApproveTokenHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldApproveToken, v))
}

// ApproveTokenHasSuffix applies the HasSuffix predicate on the "approve_token" field.
func ApproveTokenHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldApproveToken, v))
}

// ApproveTokenIsNil applies the IsNil predicate on the "approve_token" field.
func ApproveTokenIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldApproveToken))
}

// ApproveTokenNotNil applies the NotNil predicate on the "approve_token" field.
func ApproveTokenNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldApproveToken))
}

// ApproveTokenEqualFold applies the EqualFold predicate on the "approve_token" field.
func ApproveTokenEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldApproveToken, v))
}

// ApproveTokenContainsFold applies the ContainsFold predicate on the "approve_token" field.
func ApproveTokenContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldApproveToken, v))
}

// PageIDEQ applies the EQ predicate on the "page_id" field.
func PageIDEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldPageID, v))
}

// PageIDNEQ applies the NEQ predicate on the "page_id" field.
func PageIDNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldPageID, v))
}

// PageIDIn applies the In predicate on the "page_id" field.
func PageIDIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldPageID, vs...))
}

// PageIDNotIn applies the NotIn predicate on the "page_id" field.
func PageIDNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldPageID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUserID, vs...))
}

// ReplyToIDEQ applies the EQ predicate on the "reply_to_id" field.
func ReplyToIDEQ(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldReplyToID, v))
}

// ReplyToIDNEQ applies the NEQ predicate on the "reply_to_id" field.
func ReplyToIDNEQ(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldReplyToID, v))
}

// ReplyToIDIn applies the In predicate on the "reply_to_id" field.
func ReplyToIDIn(vs ...xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldReplyToID, vs...))
}

// ReplyToIDNotIn applies the NotIn predicate on the "reply_to_id" field.
func ReplyToIDNotIn(vs ...xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldReplyToID, vs...))
}

// ReplyToIDGT applies the GT predicate on the "reply_to_id" field.
func ReplyToIDGT(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldReplyToID, v))
}

// ReplyToIDGTE applies the GTE predicate on the "reply_to_id" field.
func ReplyToIDGTE(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldReplyToID, v))
}

// ReplyToIDLT applies the LT predicate on the "reply_to_id" field.
func ReplyToIDLT(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldReplyToID, v))
}

// ReplyToIDLTE applies the LTE predicate on the "reply_to_id" field.
func ReplyToIDLTE(v xid.ID) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldReplyToID, v))
}

// ReplyToIDContains applies the Contains predicate on the "reply_to_id" field.
func ReplyToIDContains(v xid.ID) predicate.Comment {
	vc := v.String()
	return predicate.Comment(sql.FieldContains(FieldReplyToID, vc))
}

// ReplyToIDHasPrefix applies the HasPrefix predicate on the "reply_to_id" field.
func ReplyToIDHasPrefix(v xid.ID) predicate.Comment {
	vc := v.String()
	return predicate.Comment(sql.FieldHasPrefix(FieldReplyToID, vc))
}

// ReplyToIDHasSuffix applies the HasSuffix predicate on the "reply_to_id" field.
func ReplyToIDHasSuffix(v xid.ID) predicate.Comment {
	vc := v.String()
	return predicate.Comment(sql.FieldHasSuffix(FieldReplyToID, vc))
}

// ReplyToIDIsNil applies the IsNil predicate on the "reply_to_id" field.
func ReplyToIDIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldReplyToID))
}

// ReplyToIDNotNil applies the NotNil predicate on the "reply_to_id" field.
func ReplyToIDNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldReplyToID))
}

// ReplyToIDEqualFold applies the EqualFold predicate on the "reply_to_id" field.
func ReplyToIDEqualFold(v xid.ID) predicate.Comment {
	vc := v.String()
	return predicate.Comment(sql.FieldEqualFold(FieldReplyToID, vc))
}

// ReplyToIDContainsFold applies the ContainsFold predicate on the "reply_to_id" field.
func ReplyToIDContainsFold(v xid.ID) predicate.Comment {
	vc := v.String()
	return predicate.Comment(sql.FieldContainsFold(FieldReplyToID, vc))
}

// HasPage applies the HasEdge predicate on the "page" edge.
func HasPage() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PageTable, PageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPageWith applies the HasEdge predicate on the "page" edge with a given conditions (other predicates).
func HasPageWith(preds ...predicate.Page) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newPageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplyTo applies the HasEdge predicate on the "reply_to" edge.
func HasReplyTo() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReplyToTable, ReplyToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReplyToWith applies the HasEdge predicate on the "reply_to" edge with a given conditions (other predicates).
func HasReplyToWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newReplyToStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplies applies the HasEdge predicate on the "replies" edge.
func HasReplies() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
func HasRepliesWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newRepliesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}
