package controller

import (
	"net/http"

	"gocode/bubble-demo/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(ctx *gin.Context) {
	//前端页面填写待办事项 点击提交 发请求到这里
	//1.从请求中把数据拿出来
	//2.存入数据库
	var todo models.Todo
	ctx.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"errir": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}

func GetTodoList(ctx *gin.Context) {
	//查询todo这个表里的所有数据
	todoList, err := models.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}

}

func UpdateATodo(ctx *gin.Context) {
	var todo models.Todo
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.GetATodo(id, &todo)
	ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if err = models.UpdateATodo(&todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
