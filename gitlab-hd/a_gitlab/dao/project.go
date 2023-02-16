package dao

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gitlab-hd/config/db"
	"gitlab-hd/a_gitlab/model"
)

//创建需要实现的接口
type IPodDao interface {
	//根据ID查找数据
	FindProjectByID(int64) (*model.Project, error)
	//查找pod所有数据
	//FindAll() ([]model.Pod,error)
}

type Projectdao struct{}

func (p *Projectdao) FindProjectByID(id int) (*model.Project, error) {
	project := &model.Project{}
	return project, db.GORM.First(project, id).Error
}

//创建pod
func (p *Projectdao) CreateProject(project *model.Project) (int64, error) {
	tx := db.GORM.Create(&project)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return project.ID, errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return project.ID, nil
}
