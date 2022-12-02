package serializer

import "todo-list/model"

type User struct {
	ID       uint   `json:"id" form:"id"`               // 用户ID
	Username string `json:"username" form:"user_name"`  // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"` // 创建
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}
