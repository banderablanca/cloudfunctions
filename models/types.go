package models

import "time"

// EventComment is the payload of a Firestore event.
type EventComment struct {
	OldValue   ValueComment `json:"oldValue"`
	Value      ValueComment `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// ValueComment holds Firestore fields.
type ValueComment struct {
	CreateTime time.Time        `json:"createTime"`
	Fields     CommentFirestore `json:"fields"`
	Name       string           `json:"name"`
	UpdateTime time.Time        `json:"updateTime"`
}

// EventFlag is the payload of a Firestore event.
type EventFlag struct {
	OldValue   ValueFlag `json:"oldValue"`
	Value      ValueFlag `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// ValueFlag holds Firestore fields.
type ValueFlag struct {
	CreateTime time.Time     `json:"createTime"`
	Fields     FlagFirestore `json:"fields"`
	Name       string        `json:"name"`
	UpdateTime time.Time     `json:"updateTime"`
}