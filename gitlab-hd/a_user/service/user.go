package service

import (
	"gitlab-hd/a_user/dao"
	"gitlab-hd/a_user/model"
)

type IUserService interface {
	SelectById(int) (*model.User, error)
	//GetGitlabProject() ([]model.Project, error)
}

func NewUserService(userdao dao.IUserdao) IUserService {
	return &UserService{
		userdao,
	}
}

type UserService struct {
	userdao dao.IUserdao
}

//单个ID查找
func (u *UserService) SelectById(id int) (user *model.User, err error) {
	user, err = u.userdao.SelectById(id)
	return user, err
}

