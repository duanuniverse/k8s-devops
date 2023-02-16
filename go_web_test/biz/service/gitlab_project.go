package service

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"go_web_test/biz/dao"
)

type ProjectService interface {
	Project(*dao.Project) (int64, error)
	GetGitlabProject() ([]*dao.Project, error)
}

type ProjectServiceImpl struct {
}

func (ProjectServiceImpl) Project(project *dao.Project) (int64, error) {
	projects := &dao.Project{}
	projects.CreateProject(project)
	return projects.ID, nil
}


func (ProjectServiceImpl) GetGitlabProject() (projectList []*dao.Project, err error) {
	token := viper.GetString("gitlab.token")
	url := viper.GetString("gitlab.url")

	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		logger.Error(errors.New("获取gitlab client失败, " + err.Error()))
		return nil, errors.New("获取gitlab client失败, " + err.Error())
	}

	for i := 1; i < 100; i++ {
		projects, resp, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page: i,
				PerPage: 100,
			},
		})
		//fmt.Println(resp.TotalItems) //获取项目总数
		if err != nil {
			logger.Error(errors.New("获取项目异常, " + err.Error()))
			return nil, errors.New("获取项目异常, " + err.Error())
		}
		//如果i大于查询总页数就介绍循环
		if i > resp.TotalPages {
			break
		}

		for  _, v := range projects {

			projectObj := &dao.Project{
				ProjectName: v.Name,
				Description: v.Description,
				WebURL:      v.WebURL,
				GroupName:   v.Namespace.Name,
			}
			projectList=append(projectList, projectObj)
			//fmt.Println(*projectObj)
		}
	}
	//fmt.Println(projectList[1].(project))
	//cells[i].(*podCell)
	return projectList,nil
}