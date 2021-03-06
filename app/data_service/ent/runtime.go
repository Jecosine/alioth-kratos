// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent/announcement"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/judgerecord"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/problem"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/schema"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/tag"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/team"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/todo"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	announcementFields := schema.Announcement{}.Fields()
	_ = announcementFields
	// announcementDescTitle is the schema descriptor for title field.
	announcementDescTitle := announcementFields[1].Descriptor()
	// announcement.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	announcement.TitleValidator = announcementDescTitle.Validators[0].(func(string) error)
	// announcementDescCreatedTime is the schema descriptor for createdTime field.
	announcementDescCreatedTime := announcementFields[3].Descriptor()
	// announcement.DefaultCreatedTime holds the default value on creation for the createdTime field.
	announcement.DefaultCreatedTime = announcementDescCreatedTime.Default.(time.Time)
	// announcementDescModifiedTime is the schema descriptor for modifiedTime field.
	announcementDescModifiedTime := announcementFields[4].Descriptor()
	// announcement.DefaultModifiedTime holds the default value on creation for the modifiedTime field.
	announcement.DefaultModifiedTime = announcementDescModifiedTime.Default.(time.Time)
	judgerecordFields := schema.JudgeRecord{}.Fields()
	_ = judgerecordFields
	// judgerecordDescStatus is the schema descriptor for status field.
	judgerecordDescStatus := judgerecordFields[5].Descriptor()
	// judgerecord.DefaultStatus holds the default value on creation for the status field.
	judgerecord.DefaultStatus = judgerecordDescStatus.Default.(int64)
	problemFields := schema.Problem{}.Fields()
	_ = problemFields
	// problemDescTitle is the schema descriptor for title field.
	problemDescTitle := problemFields[1].Descriptor()
	// problem.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	problem.TitleValidator = problemDescTitle.Validators[0].(func(string) error)
	// problemDescCreatedTime is the schema descriptor for created_time field.
	problemDescCreatedTime := problemFields[3].Descriptor()
	// problem.DefaultCreatedTime holds the default value on creation for the created_time field.
	problem.DefaultCreatedTime = problemDescCreatedTime.Default.(time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescCreatedTime is the schema descriptor for created_time field.
	tagDescCreatedTime := tagFields[2].Descriptor()
	// tag.DefaultCreatedTime holds the default value on creation for the created_time field.
	tag.DefaultCreatedTime = tagDescCreatedTime.Default.(time.Time)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreatedTime is the schema descriptor for created_time field.
	teamDescCreatedTime := teamFields[3].Descriptor()
	// team.DefaultCreatedTime holds the default value on creation for the created_time field.
	team.DefaultCreatedTime = teamDescCreatedTime.Default.(time.Time)
	// teamDescPrivate is the schema descriptor for private field.
	teamDescPrivate := teamFields[4].Descriptor()
	// team.DefaultPrivate holds the default value on creation for the private field.
	team.DefaultPrivate = teamDescPrivate.Default.(bool)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[1].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescDone is the schema descriptor for done field.
	todoDescDone := todoFields[2].Descriptor()
	// todo.DefaultDone holds the default value on creation for the done field.
	todo.DefaultDone = todoDescDone.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[1].Descriptor()
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[4].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescCreatedTime is the schema descriptor for created_time field.
	userDescCreatedTime := userFields[5].Descriptor()
	// user.DefaultCreatedTime holds the default value on creation for the created_time field.
	user.DefaultCreatedTime = userDescCreatedTime.Default.(time.Time)
}
