package main

import (
	"gocode/bubble-demo/dao"
	"gocode/bubble-demo/models"
	"gocode/bubble-demo/routers"
)

func main() {
	//创建数据库
	//sql:CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()

}
