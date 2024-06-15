// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/rs/xid"
	"github.com/supersupersimple/comment/ent/comment"
	"github.com/supersupersimple/comment/ent/conf"
	"github.com/supersupersimple/comment/ent/page"
	"github.com/supersupersimple/comment/ent/schema"
	"github.com/supersupersimple/comment/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[3].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[4].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// commentDescDepth is the schema descriptor for depth field.
	commentDescDepth := commentFields[5].Descriptor()
	// comment.DefaultDepth holds the default value on creation for the depth field.
	comment.DefaultDepth = commentDescDepth.Default.(int)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.DefaultID holds the default value on creation for the id field.
	comment.DefaultID = commentDescID.Default.(func() xid.ID)
	confFields := schema.Conf{}.Fields()
	_ = confFields
	// confDescLimitPerBatch is the schema descriptor for limit_per_batch field.
	confDescLimitPerBatch := confFields[3].Descriptor()
	// conf.DefaultLimitPerBatch holds the default value on creation for the limit_per_batch field.
	conf.DefaultLimitPerBatch = confDescLimitPerBatch.Default.(int)
	// confDescMaxLoopDepth is the schema descriptor for max_loop_depth field.
	confDescMaxLoopDepth := confFields[4].Descriptor()
	// conf.DefaultMaxLoopDepth holds the default value on creation for the max_loop_depth field.
	conf.DefaultMaxLoopDepth = confDescMaxLoopDepth.Default.(int)
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescCreatedAt is the schema descriptor for created_at field.
	pageDescCreatedAt := pageFields[3].Descriptor()
	// page.DefaultCreatedAt holds the default value on creation for the created_at field.
	page.DefaultCreatedAt = pageDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[3].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}