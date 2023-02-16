package service

import "go_web_test/biz/dao"

type UserService interface {
	User(userId int) *dao.User
}

type UserServiceImpl struct {
}

func (UserServiceImpl) User(userId int) *dao.User {
	user := &dao.User{}
	user.SelectById(userId)
	return user
}
