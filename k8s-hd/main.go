package main

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"k8s-hd/config"
	"k8s-hd/db"
	"k8s-hd/model"
	"log"
)

func main() {
	//初始化数据库
	db.Init()

	git, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.Url))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	log.Println("开始获取项目信息")
	projectList := []interface{}{}

	for i := 1; i < 100; i++ {
		projects, resp, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page: i,
				PerPage: 100,
			},
		})
		//fmt.Println(resp.TotalItems) //获取项目总数
		if err != nil {
			fmt.Printf("获取项目信息异常:%v", err)
			panic(err)
		}
		//如果i大于查询总页数就介绍循环
		if i > resp.TotalPages {
			break
		}

		for  _, v := range projects {
			projectObj := &model.Project{
				Description: v.Description,
				ProjectName: v.Name,
				WebURL:      v.WebURL,
				GroupName:   v.Namespace.Name,
			}
			projectList=append(projectList, *projectObj)
			fmt.Println(*projectObj)
		}
	}
	//fmt.Println(projectList[1].(project))
	//cells[i].(*podCell)
}
