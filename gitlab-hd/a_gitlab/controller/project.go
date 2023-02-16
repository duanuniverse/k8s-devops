package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab-hd/a_gitlab/service"
	"gitlab-hd/dto"
	"net/http"
	"strconv"
)

var Project ProjectHandler

type ProjectHandler struct {
	//注意这里的类型实 IProjectService 接口类型
	ProjectService service.IProjectService
}

//查询单个信息
func (e *ProjectHandler) FindProjectByID(ctx *gin.Context) {

	id := ctx.Param("id")
	PId, _ := strconv.Atoi(id)
	
	project, _ := e.ProjectService.FindProjectByID(int64(PId))

	ctx.JSON(http.StatusOK, dto.Ok(project))
}

/*//查询所有项目
func (*Projectct) GetAllGitProject(ctx *gin.Context) {
	data, err := service.Projectservice.GetGitlabProject()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Ok(data))
}

func TestFunc(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "test成功",
		"data": nil,
	})
}*/
