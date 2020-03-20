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
	managedProjects.Projects = models.GetJoinProjects(userID)
	//返回json数组
	c.JSON(http.StatusOK, managedProjects)
}
