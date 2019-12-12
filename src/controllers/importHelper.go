package controllers

import (
	"Kcoin-Golang/src/models"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ImportProject(url string,cover_url string)error{

	//首先将string类型的currentUserId转成Int型
	currentUserId_int,err:=strconv.Atoi(currentUserId)
	//检查地址是否合法
	err = CheckGithubRepoUrl(currentUserId, url)
	if err!=nil{
		log.Fatal("url is illegal",err)
		return err
	}
	//解析已经合法的地址中的用户名和仓库名
	userName, repoName, _:=models.ParseGithubHTTPSUrl(url)
	//使用用户名和仓库名获取项目全部contributor
	userslist_string:=models.GetContributors(userName,repoName)
	users:=strings.Split(userslist_string," ")
	//将当前登录用户注册到webhook中，
	registerGithubWebhooks(currentUserId, repoName)

	//查询数据库，看看所有的contributor哪个已经加入到数据库了，如果已经加入数据库，就把这些人放到k_user_in_project这个表里
	var alreadyIn []string//已经加入kcoin的contributor
	var notIn []string//还没加入kcoin的contributor
	for _,singleUser:=range users{
		//遍历全部contributor列表
		var k_user_id int
		var project_id int
		res, err := models.FindUserByUsername(userName)
		if err!=nil{
			log.Fatal("when query a single user in k_user table,an error occured,",err)
		}else{
			num, _ :=res.RowsAffected()
			if num==0{//如果没查到
				notIn= append(notIn, singleUser)
				//通过github API获取这个用户的git id
				singleUser_git_id:=GetGithubId(singleUser)
				//插入项目到数据库
				_, _ = models.InsertProject(repoName, url, cover_url)
				//根据根据项目名查K_Project表获取project_id
				project_id,err=models.GetProjectidByRepoName(repoName)
				if err!=nil{
					log.Fatal("when query project_id in K_Project ,error occured:",err)
				}
				//插入到临时用户表
				//TODO :（其实也不是TODO，是一个提醒）现在这个地方应该是插入不成功的，因为temporary_user_id还不是自动生成的，等待数据库方面将其设置为自增就行了。
				res,err:=models.InsertIntoKTemporaryUser(currentUserId_int,singleUser_git_id,singleUser,project_id)
				if err!=nil{
					num, _ :=res.RowsAffected()
					fmt.Println(num)
				}else{
					log.Fatal("when insert into k_temporary_user,error occured,",err)
				}
			}else{//查到了，这个用户已经在项目了
				alreadyIn=append(alreadyIn,singleUser)
				//然后把这个用户加到k_user_in_project中
				//首先根据user_name获取k_user_id
				k_user_id,err=models.GetUseridByUsername(userName)
				if err!=nil{
					log.Fatal("when query k_user_id in K_User,error occured:",err)
				}
				//然后根据根据项目名查K_Project表获取project_id
				project_id,err=models.GetProjectidByRepoName(repoName)
				if err!=nil{
					log.Fatal("when query project_id in K_Project ,error occured:",err)
				}
				//最后两个id都有了，插入K_user_in_project
				res,err =models.InsertIntoKUserInProject(project_id,k_user_id)
				if err==nil{
					nums, _ :=res.RowsAffected()
					fmt.Println("affected rows has",nums)
					id,_:=res.LastInsertId()
					fmt.Println("last insert id is ",id)
				}else{
					log.Fatal("when insert to K_User_in_Project,error occured,",err)
				}
			}
		}
	}
	//对所有没有加入的人给他们发个邮件
	_, err = models.SendEMailToPotentialUsers(notIn)
	return err
}

