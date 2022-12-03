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

type DeleteTaskService struct {
}

// 更新任务的服务
type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

// 搜索任务的服务
type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

func (s *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	code := e.SUCCESS
	err := model.DB.Where("uid=?", uid).Preload("User").
		Where("title LIKE ? OR content LIKE ?",
			"%"+s.Info+"%", "%"+s.Info+"%").Find(&tasks).Error
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
		Data:   serializer.BuildTasks(tasks),
	}
}

func (s *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = s.Content
	task.Status = s.Status
	task.Title = s.Title
	code := e.SUCCESS
	err := model.DB.Save(&task).Error
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
		Data:   "修改成功",
	}
}

func (s *DeleteTaskService) Delete(id string) serializer.Response {
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
	err = model.DB.Delete(&task).Error
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
	}
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
