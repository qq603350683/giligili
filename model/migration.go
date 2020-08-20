package model

// 执行数据迁移

func migration() {
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserToken{})
	DB.AutoMigrate(&Level{})
	DB.AutoMigrate(&UserPlan{})
}
