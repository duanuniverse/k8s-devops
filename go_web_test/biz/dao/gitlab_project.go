package dao

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"go_web_test/config/db"
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


func (p *Project) CreateProject(project *Project) (int64, error) {
	tx := db.DB.Create(&project)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return project.ID, errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return project.ID, nil
}