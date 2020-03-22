package controller

import (
	"Kcoin-Golang/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO 从cookie中获取用户id
func GetUserId() (string, error) {
	return "149", nil
}

//获取当前用户参与的项目
func GetJoinProjects(c *gin.Context) {
	var joinedProjects models.ProjectInfo
	var userID string
	var err error
	//获取用户ID并进行错误处理
	//TODO 错误处理
	if userID, err = GetUserId(); err != nil {
		fmt.Println(err.Error())
	}
	//获取当前用户参与的项目
	joinedProjects.Projects = models.GetJoinProjects(userID)
	//返回json数组
	c.JSON(http.StatusOK, joinedProjects)
}

//获取当前用户管理的项目
func GetManageProjects(c *gin.Context) {
	var managedProjects models.ProjectInfo
	var userID string
	var err error
	//获取用户ID并进行错误处理
	//TODO 错误处理
	if userID, err = GetUserId(); err != nil {
		fmt.Println(err.Error())
	}
	//获取当前用户管理的项目
	managedProjects.Projects = models.GetManageProjects(userID)
	//返回json数组
	c.JSON(http.StatusOK, managedProjects)
}

func AddProject(c *gin.Context) {
	// TODO: 判断是否是已注册用户
	var project *models.ProjectDetail
	// 解析传入的json数据，并绑定到project上。若失败，将返回错误并在http头部写入400状态码
	c.BindJSON(&project)    
	// TODO: 需要判断Url是否合法
	if urlLegal := true; !urlLegal {
		c.JSON(http.StatusOK, gin.H{"result":"项目url不合法，请检查"})
		return
	}
	// TODO: 需要判断用户是否有权限将项目导入平台
	if authorized := true; !authorized {
		c.JSON(http.StatusOK, gin.H{"result":"抱歉，您没有导入项目的权限"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": project.GithubUrl,
		"name": project.ProjectName,
	})
/*
	code := models.AddProject(project)
	var result string
	if code == 0 {
		// url已经存在，说明项目已经在平台上，进行相应处理
		result = "项目已在平台中"
	} else if code == -1 {
		result = "项目导入失败"
	} else {
		result = "项目导入成功"
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
*/	
}