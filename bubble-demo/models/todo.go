package models

import "gocode/bubble-demo/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 这个Model的增删改查操作都放在这里

// CreateATodo 创建todo
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	//查询todo这个表里的所有数据
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string, todo *Todo) (err error) {
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return err
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
