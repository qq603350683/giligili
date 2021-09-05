package service

import (
	"giligili/model"
	"time"
)

type DailyResult struct {
	TodayRegister int `json:"today_register"`
	TodayPassLevel int `json:"today_pass_level"`
	YesterdayRegister int `json:"yesterday_register"`
	YesterdayPassLevel int `json:"yesterday_pass_level"`
}

func NewDailyResult() *DailyResult {
	return new(DailyResult)
}

func GetDailyData() *DailyResult {
	dailyResult := NewDailyResult()

	user := model.NewUser()

	dailyResult.TodayRegister = user.GetRegisterCount(time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")+" 23:59:59")
	dailyResult.TodayPassLevel = user.GetPassLevelCount(time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")+" 23:59:59")
	dailyResult.YesterdayRegister = user.GetRegisterCount(time.Now().AddDate(0, 0, -1).Format("2006-01-02"), time.Now().AddDate(0, 0, -1).Format("2006-01-02")+" 23:59:59")
	dailyResult.YesterdayPassLevel = user.GetPassLevelCount(time.Now().AddDate(0, 0, -1).Format("2006-01-02"), time.Now().AddDate(0, 0, -1).Format("2006-01-02")+" 23:59:59")

	return dailyResult
}
