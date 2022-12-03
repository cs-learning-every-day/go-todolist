package api

import (
	"todo-list/pkg/e"
	util "todo-list/pkg/utils"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

func SearchTasks(c *gin.Context) {
	s := service.SearchTaskService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Search(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UpdateTask(c *gin.Context) {
	s := service.UpdateTaskService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func DeleteTask(c *gin.Context) {
	s := service.DeleteTaskService{}
	res := s.Delete(c.Param("id"))
	c.JSON(e.SUCCESS, res)
}

func ListTasks(c *gin.Context) {
	s := service.ListTaskService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.List(claim.Id)
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func ShowTask(c *gin.Context) {
	s := service.ShowTaskService{}
	res := s.Show(c.Param("id"))
	c.JSON(e.SUCCESS, res)
}

func CreateTask(c *gin.Context) {
	s := service.CreateTaskSertice{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create(claim.Id)
		c.JSON(e.SUCCESS, res)
	} else {
		c.JSON(e.InvalidParams, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
