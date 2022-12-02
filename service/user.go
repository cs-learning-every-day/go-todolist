package service

import (
	"errors"
	"todo-list/model"
	"todo-list/pkg/e"
	util "todo-list/pkg/utils"
	"todo-list/serializer"

	"gorm.io/gorm"
)

type UserService struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (us *UserService) Login() *serializer.Response {
	var user model.User
	code := e.SUCCESS
	// check user exist
	if err := model.DB.Where("username=?", us.Username).First(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = e.ErrorNotExistUser
			return &serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !user.CheckPassword(us.Password) {
		code = e.ErrorNotCompare
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, user.Username, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

func (us *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64

	model.DB.Model(&model.User{}).
		Where("username=?", us.Username).
		First(&user).Count(&count)
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Username = us.Username
	if err := user.SetPassword(us.Password); err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
