package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_test/biz/dto"
	"go_web_test/biz/service"
	"net/http"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

func ProjectApi(router *gin.Engine) {

	/*userHandler := UserHandler{
		userService: &service.UserServiceImpl{},
	}*/
	projectHandler := ProjectHandler{
		projectService: service.ProjectServiceImpl{},
	}

	userGroup := router.Group("project/")
	{
		userGroup.GET("/getlist", projectHandler.GetGitlabProject)
	}
}

func (projectHandler ProjectHandler) GetGitlabProject(c *gin.Context) {
	projectList, _ := projectHandler.projectService.GetGitlabProject()
	c.JSON(http.StatusOK, dto.Ok(projectList))
}
