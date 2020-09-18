package model

// 执行数据迁移

func migration() {
	//DB.AutoMigrate(&Video{})
	//DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Backpack{})
	//DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	set := "ENGINE=InnoDB"

	if DB.HasTable(&User{}) {
		DB.AutoMigrate(&User{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&User{})
	}

	if DB.HasTable(&UserToken{}) {
		DB.AutoMigrate(&UserToken{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&UserToken{})
	}

	if DB.HasTable(&Level{}) {
		DB.AutoMigrate(&Level{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&Level{})
	}

	if DB.HasTable(&UserPlan{}) {
		DB.AutoMigrate(&UserPlan{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&UserPlan{})
	}

	if DB.HasTable(&SignIn{}) {
		DB.AutoMigrate(&SignIn{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&SignIn{})
	}

	if DB.HasTable(&Backpack{}) {
		DB.AutoMigrate(&Backpack{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&Backpack{})
	}

	if DB.HasTable(&ForwardPrize{}) {
		DB.AutoMigrate(&ForwardPrize{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&ForwardPrize{})
	}

	if DB.HasTable(&Prop{}) {
		DB.AutoMigrate(&Prop{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&Prop{})
	}

	if DB.HasTable(&SignInPrize{}) {
		DB.AutoMigrate(&SignInPrize{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&SignInPrize{})
	}

	if DB.HasTable(&UserSignInPrize{}) {
		DB.AutoMigrate(&UserSignInPrize{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&UserSignInPrize{})
	}

	if DB.HasTable(&Store{}) {
		DB.AutoMigrate(&Store{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&Store{})
	}

	if DB.HasTable(&UserPassLevel{}) {
		DB.AutoMigrate(&UserPassLevel{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&UserPassLevel{})
	}

	if DB.HasTable(&Plan{}) {
		DB.AutoMigrate(&Plan{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&Plan{})
	}

	if DB.HasTable(&WechatUser{}) {
		DB.AutoMigrate(&WechatUser{})
	} else {
		DB.Set("gorm:table_options", set).CreateTable(&WechatUser{})
	}

	//DB.AutoMigrate(&User{})
	//DB.AutoMigrate(&UserToken{})
	//DB.AutoMigrate(&Level{})
	//DB.AutoMigrate(&UserPlan{})
	//DB.AutoMigrate(&SignIn{})
	//DB.AutoMigrate(&Backpack{})
	//DB.AutoMigrate(&ForwardPrize{})
	//DB.AutoMigrate(&Prop{})
	//DB.AutoMigrate(&SignInPrize{})
	//DB.AutoMigrate(&UserSignInPrize{})
	//DB.AutoMigrate(&Store{})
	//DB.AutoMigrate(&UserPassLevel{})
	//DB.AutoMigrate(&Plan{})
	//DB.AutoMigrate(&WechatUser{})
}
