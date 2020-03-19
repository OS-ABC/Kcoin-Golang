package controller

import (
	"Kcoin-Golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO 从cookie中获取用户id
func GetUserId(){
}

//获取当前用户参与的项目
func GetJoinProjects(c *gin.Context){
	var joinedProjects models.ProjectInfo
	//获取用户ID
	var userID string
	//userID = GetUserId()
	userID = "149"
	//获取当前用户参与的项目
	joinedProjects.Projects = models.GetJoinProjects(userID)
	//返回json数组
	c.JSON(http.StatusOK,joinedProjects)
}

//获取当前用户管理的项目
func GetManageProjects(c *gin.Context){
	var managedProjects models.ProjectInfo
	//获取用户ID
	var userID string
	//userID = GetUserId()
	userID = "149"
	//获取当前用户管理的项目
	managedProjects.Projects = models.GetJoinProjects(userID)
	//返回json数组
	c.JSON(http.StatusOK,managedProjects)
}
