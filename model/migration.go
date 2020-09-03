package model

// 执行数据迁移

func migration() {
	//DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserToken{})
	DB.AutoMigrate(&Level{})
	DB.AutoMigrate(&UserPlan{})
	DB.AutoMigrate(&SignIn{})
	DB.AutoMigrate(&Backpack{})
	DB.AutoMigrate(&ForwardPrize{})
	DB.AutoMigrate(&Prop{})
	DB.AutoMigrate(&SignInPrize{})
	DB.AutoMigrate(&UserSignInPrize{})
	DB.AutoMigrate(&Store{})
	DB.AutoMigrate(&UserPassLevel{})
}
