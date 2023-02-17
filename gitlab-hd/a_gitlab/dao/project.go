package dao

import (
	"github.com/jinzhu/gorm"
	"gitlab-hd/a_gitlab/model"
)

//创建需要实现的接口
type IProjectdao interface {
	//根据ID查找数据
	FindProjectByID(int) (*model.Project, error)
	//查找pod所有数据
	//FindAll() ([]model.Pod,error)
}

func NewProjectdao(db *gorm.DB) IProjectdao {
	return &Projectdao{
		db,
	}
}

type Projectdao struct{
	db *gorm.DB
}

func (p *Projectdao) FindProjectByID(id int) (*model.Project, error) {
	project := &model.Project{}
	return project, p.db.First(project, id).Error
	//return project, db.DB.First(project, id).Error
}

//创建pod
/*func (p *Projectdao) CreateProject(project *model.Project) (int64, error) {
	tx := db.GORM.Create(&project)
	if tx.Error != nil {
		logger.Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return project.ID, errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return project.ID, nil
}*/
