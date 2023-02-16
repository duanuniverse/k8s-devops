package model

import (
	"time"
)

type Project struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	ProjectName		string      `json:"project_name"`
	Description		string      `json:"description"`
	WebURL			string      `json:"web_url"`
	GroupName		string      `json:"group_name"`
}

func(*Project) TableName() string {
	return "git_project"
}