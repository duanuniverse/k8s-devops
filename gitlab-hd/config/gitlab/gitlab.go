package gin

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

// InitGitlabClient 初始化gitlab
func InitGitlabClient() (git *gitlab.Client, err error) {
	token := viper.GetString("gitlab.token")
	url := viper.GetString("gitlab.url")
	git, err = gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		logger.Error(errors.New("获取gitlab client失败, " + err.Error()))
		return nil, errors.New("获取gitlab client失败, " + err.Error())
	}
	return git, err
}


