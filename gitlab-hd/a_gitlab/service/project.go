package service

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
	"gitlab-hd/a_gitlab/dao"
	"gitlab-hd/a_gitlab/model"
	mygitlab "gitlab-hd/config/gitlab"
)

type IProjectService interface {
	FindProjectByID(int64) (*model.Project, error)
	//GetGitlabProject() ([]model.Project, error)
}

type ProjectService struct {
	ProjectDao dao.IPodDao
}

//单个ID查找
func (p *ProjectService) FindProjectByID(id int64) (*model.Project, error) {
	return p.ProjectDao.FindProjectByID(id)
}

func (u *ProjectService) GetGitlabProject() (projectList []model.Project, err error) {

	git, err := mygitlab.InitGitlabClient()
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

			projectObj := &model.Project{
				ProjectName: v.Name,
				Description: v.Description,
				WebURL:      v.WebURL,
				GroupName:   v.Namespace.Name,
			}
			projectList=append(projectList, *projectObj)
			//fmt.Println(*projectObj)
		}
	}
	return projectList,nil
}