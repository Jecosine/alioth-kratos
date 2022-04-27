// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/Jecosine/alioth-kratos/app/data_service/ent"
)

type JudgeReply struct {
	ID           string    `json:"id"`
	User         *ent.User `json:"user"`
	JudgeTime    time.Time `json:"judgeTime"`
	FinishedTime time.Time `json:"finishedTime"`
	TimeCost     int       `json:"timeCost"`
	MemoryCost   int       `json:"memoryCost"`
	Status       int       `json:"status"`
}

type JudgeRequest struct {
	ID          string       `json:"id"`
	User        *ent.User    `json:"user"`
	CreatedTime time.Time    `json:"createdTime"`
	SourceCode  string       `json:"source_code"`
	Problem     *ent.Problem `json:"problem"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type PingInput struct {
	Msg string `json:"msg"`
}

type Privilege struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Scope      *RoleScope   `json:"scope"`
	Privileges []*Privilege `json:"privileges"`
}

type Task struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	TaskType   *TaskType   `json:"taskType"`
	Content    *string     `json:"content"`
	Assignees  []*ent.User `json:"assignees"`
	Status     TaskStatus  `json:"status"`
	ParentList *TaskList   `json:"parentList"`
}

type TaskList struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Assignees    []*ent.User `json:"assignees"`
	Deadline     time.Time   `json:"deadline"`
	Status       TaskStatus  `json:"status"`
	RelyTaskList *TaskList   `json:"relyTaskList"`
	Tasks        []*Task     `json:"tasks"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RoleScope string

const (
	RoleScopeGlobal RoleScope = "GLOBAL"
	RoleScopeTeam   RoleScope = "TEAM"
)

var AllRoleScope = []RoleScope{
	RoleScopeGlobal,
	RoleScopeTeam,
}

func (e RoleScope) IsValid() bool {
	switch e {
	case RoleScopeGlobal, RoleScopeTeam:
		return true
	}
	return false
}

func (e RoleScope) String() string {
	return string(e)
}

func (e *RoleScope) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleScope(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleScope", str)
	}
	return nil
}

func (e RoleScope) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RoleType string

const (
	RoleTypeAdmin     RoleType = "ADMIN"
	RoleTypeTeamAdmin RoleType = "TEAM_ADMIN"
	RoleTypeUser      RoleType = "USER"
)

var AllRoleType = []RoleType{
	RoleTypeAdmin,
	RoleTypeTeamAdmin,
	RoleTypeUser,
}

func (e RoleType) IsValid() bool {
	switch e {
	case RoleTypeAdmin, RoleTypeTeamAdmin, RoleTypeUser:
		return true
	}
	return false
}

func (e RoleType) String() string {
	return string(e)
}

func (e *RoleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleType", str)
	}
	return nil
}

func (e RoleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusExpired  TaskStatus = "Expired"
	TaskStatusFinished TaskStatus = "Finished"
	TaskStatusWaiting  TaskStatus = "Waiting"
	TaskStatusCanceled TaskStatus = "Canceled"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusExpired,
	TaskStatusFinished,
	TaskStatusWaiting,
	TaskStatusCanceled,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusExpired, TaskStatusFinished, TaskStatusWaiting, TaskStatusCanceled:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskType string

const (
	TaskTypeProblem    TaskType = "Problem"
	TaskTypeCustom     TaskType = "Custom"
	TaskTypeMeeting    TaskType = "Meeting"
	TaskTypeSubmission TaskType = "Submission"
)

var AllTaskType = []TaskType{
	TaskTypeProblem,
	TaskTypeCustom,
	TaskTypeMeeting,
	TaskTypeSubmission,
}

func (e TaskType) IsValid() bool {
	switch e {
	case TaskTypeProblem, TaskTypeCustom, TaskTypeMeeting, TaskTypeSubmission:
		return true
	}
	return false
}

func (e TaskType) String() string {
	return string(e)
}

func (e *TaskType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskType", str)
	}
	return nil
}

func (e TaskType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}