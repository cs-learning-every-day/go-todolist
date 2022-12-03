package service

import (
	"time"
	"todo-list/model"
	"todo-list/pkg/e"
	util "todo-list/pkg/utils"
	"todo-list/serializer"
)

type CreateTaskSertice struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

type ShowTaskService struct {
}

type ListTaskService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

func (s *ListTaskService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	if s.Limit == 0 {
		s.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").
		Where("uid = ?", id).Count(&total).
		Limit(s.Limit).Offset((s.Start - 1) * s.Limit).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))

}

func (s *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, id).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	task.AddView()
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

func (s *CreateTaskSertice) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     s.Title,
		Content:   s.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}

	code := e.SUCCESS
	err := model.DB.Create(&task).Error
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTask(task),
	}
}
