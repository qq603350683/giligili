package util

import (
	"log"
	"time"
)

const (
	DATE = "2006-01-02"
	DATETIME = "2006-01-02 15:04:05"
)

// time.Time 格式转化为字符串格式
func ToDatetime(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}

// 字符串转化为 time.Time
func ToTime(str string) time.Time {
	local, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, local)
	if err != nil {
		panic("ToTime: '" + str + "' 转化为time.Time格式错误")
	}

	return t;
}

// 判断是否日期格式
func IsDatetime(str string) bool {
	_, err := time.Parse(DATE, str)
	if err != nil {
		log.Printf("日期格式(%s)不正确", str)
		return false
	}

	return true
}