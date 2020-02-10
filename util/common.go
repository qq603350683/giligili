package util

import (
	"errors"
	"html"
	"strconv"

	//"strconv"
)

// 转化为数字类型
func ToInt(str string) (int, error) {
	if str == "" {
		return 0, errors.New("请勿输入空字符串")
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, errors.New("(" + str + ")字符串转换数字失败")
	}

	return i, nil
}

// 字符串过滤
func Filter(str string) string {
	str = XssFilter(str)

	// 敏感词
	sensitive_works := [1]string{"垃圾"}

	return sensitive_works[0]
}

// 过滤xss攻击
func XssFilter(str string) string {
	return html.EscapeString(str)
}

// 获取数据的类型
func GetType(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return "string", nil
	case int:
		return "int", nil
	default:
		return "", errors.New("未知类型")
	}
}

// 判断是否字符串类型
func IsString(data interface{}) bool {
	t, _ := GetType(data)
	if t == "string" {
		return true
	}

	return false
}