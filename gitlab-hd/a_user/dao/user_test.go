package dao

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gitlab-hd/config/db"
	"testing"
)

const filePath = "../../application.yaml"

func TestGetUser(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}
	db.InitDB()
	uerss := NewUserdao(db.DB)
	id := 1

	user, _ := uerss.SelectById(id)
	logger.Info(user)

	assert.NotNil(t, user)
}
