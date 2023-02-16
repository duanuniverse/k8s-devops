package dao

import "go_web_test/config/db"

type User struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:50"`
}

func (user *User) SelectById(userId int) {
	db.DB.First(&user, userId)
}
