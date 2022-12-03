package api

import (
	"todo-list/pkg/e"
	util "todo-list/pkg/utils"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

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
