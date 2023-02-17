package dao

import (
	"github.com/jinzhu/gorm"
	"gitlab-hd/a_user/model"
)

//创建需要实现的接口
type IUserdao interface {
	//根据ID查找数据
	SelectById(int) (*model.User, error)
	//查找pod所有数据
	//FindAll() ([]model.Pod,error)
}

func NewUserdao(db *gorm.DB) IUserdao {
	return &Userdao{
		db,
	}
}

type Userdao struct{
	db *gorm.DB
}

func (u *Userdao) SelectById(userId int) (*model.User, error) {
	user := &model.User{}
	return user, u.db.First(u.db.First(user, userId)).Error
}