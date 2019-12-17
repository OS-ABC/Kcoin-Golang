// TODO 把该文件与github_helper.go移动到service文件夹下, 并保证可以运行
package controllers

import (
	"Kcoin-Golang/src/models"
	"Kcoin-Golang/src/service"
	"fmt"
	"log"
	"strings"
)

func ImportProject(url string, cover_url string) error {
	fmt.Println("进入ImportProject")
	//首先将string类型的currentUserId转成Int型
	//currentUserId_int,err:=strconv.Atoi(currentUserId)
	//检查地址是否合法
	err := service.CheckGithubRepoUrl(currentUserId, url)
	if err != nil {
		log.Fatal("url is illegal", err)
		return err
	} else {
		fmt.Println("url合法")
	}
	//解析已经合法的地址中的用户名和仓库名
	userName, repoName, _ := service.ParseGithubHTTPSUrl(url)
	//使用用户名和仓库名获取项目全部contributor
	userslist_string := service.GetContributors(userName, repoName)
	users := strings.Split(userslist_string, " ")
	//将当前登录用户注册到webhook中，
	fmt.Println(currentUserId)
	service.RegisterGithubWebhooks(currentUserId, repoName)
	host_id, _ := models.GetUseridByUsername(userName)
	fmt.Println("当前登陆用户id为", host_id, "当前username为", userName)
	fmt.Println("项目中的全部contributors", users)
	//查询数据库，看看所有的contributor哪个已经加入到数据库了，如果已经加入数据库，就把这些人放到k_user_in_project这个表里
	var alreadyIn []string //已经加入kcoin的contributor
	var notIn []string     //还没加入kcoin的contributor
	var project_id int

	//插入项目到数据库
	_, _ = models.InsertProject(repoName, url, cover_url)
	//根据根据项目名查K_Project表获取project_id
	project_id, err = models.GetProjectidByRepoName(repoName)
	if err != nil {
		log.Fatal("when query project_id in K_Project ,error occured:", err)
	}

	for _, singleUser := range users {
		fmt.Println("开始for循环")
		if singleUser == "" {
			break
		}
		//遍历全部contributor列表
		var k_user_id int

		fmt.Println("正在找当前contributor:", singleUser)
		res, err := models.FindUserByUsername(singleUser)
		if err != nil {
			log.Fatal("when query a single user in k_user table,an error occured,", err)
		} else {
			num, _ := res.RowsAffected()
			if num == 0 { //如果没查到
				notIn = append(notIn, singleUser)
				//通过github API获取这个用户的git id
				singleUser_git_id := service.GetGithubId(singleUser)
				//插入到临时用户表
				//TODO :（其实也不是TODO，是一个提醒）现在这个地方应该是插入不成功的，因为temporary_user_id还不是自动生成的，等待数据库方面将其设置为自增就行了。
				res, err := models.InsertIntoKTemporaryUser(host_id, singleUser_git_id, singleUser, project_id)
				if err == nil {
					num, _ := res.RowsAffected()
					fmt.Println(num)
				} else {
					log.Fatal("when insert into k_temporary_user,error occured,", err)
				}
			} else { //查到了，这个用户已经在项目了
				alreadyIn = append(alreadyIn, singleUser)
				//然后把这个用户加到k_user_in_project中
				//首先根据user_name获取k_user_id
				k_user_id, err = models.GetUseridByUsername(singleUser)
				if err != nil {
					log.Fatal("when query k_user_id in K_User,error occured:", err)
				}
				//然后根据根据项目名查K_Project表获取project_id
				project_id, err = models.GetProjectidByRepoName(repoName)
				if err != nil {
					log.Fatal("when query project_id in K_Project ,error occured:", err)
				}
				//最后两个id都有了，插入K_user_in_project
				res, err = models.InsertIntoKUserInProject(project_id, k_user_id)
				if err == nil {
					nums, _ := res.RowsAffected()
					fmt.Println("affected rows has", nums)
					id, _ := res.LastInsertId()
					fmt.Println("last insert id is ", id)
				} else {
					log.Fatal("when insert to K_User_in_Project,error occured,", err)
				}
			}
		}
	}
	//最后判断一下项目拥有者在不在K_user_in_Projet里，不在就插入
	var IfOwnerIn int
	IfOwnerIn, err = models.FindUserInKUserInProject(host_id)
	if IfOwnerIn == 0 {
		fmt.Println("拥有者还不在数据库，插入")

		_, err = models.InsertIntoKUserInProject(project_id, host_id)
	}
	//对所有没有加入的人给他们发个邮件
	_, err = models.SendEMailToPotentialUsers(notIn)
	return err
}
