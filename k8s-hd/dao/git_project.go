package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
	"k8s-hd/db"
	"k8s-hd/model"
	"fmt"
)

//创建需要实现的接口
type IProjectDao interface {
	//创建一条Pod 数据
	CreatePod(project *model.Project) (int64,error)
}

type ProjectDao struct {

}

//创建pod
func (p *ProjectDao) CreateProject(project *model.Project) (int64, error) {
	tx := db.GORM.Create(&project)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return project.ID, errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return project.ID, nil
}