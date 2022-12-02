package api

import (
	util "todo-list/pkg/utils"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	} else {
		res := userService.Register()
		c.JSON(200, res)
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	} else {
		res := userService.Login()
		c.JSON(200, res)
	}
}
