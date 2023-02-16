package service

import (
	"k8s-hd/dao"
	"k8s-hd/model"
	"k8s.io/client-go/kubernetes"
)

type IProjectService interface {
	AddProject(project *model.Project) (int64, error)
	Get
}

func NewProjectService(projectDao dao.IProjectDao, clientset *kubernetes.Clientset) IProjectService {
	return &ProjectService{
		ProjectDao: projectDao,
	}
}

type ProjectService struct {
	ProjectDao dao.IProjectDao
}

//添加pod
func (u *ProjectService) AddProject(project *model.Project) (int64, error) {
	return u.ProjectDao.CreatePod(project)
}