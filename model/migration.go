package model

// 执行数据迁移
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{})
	if err != nil {
		return
	}
	//DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}
