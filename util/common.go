package util

import (
	"encoding/json"
	"errors"
	"html"
	"log"
	"math/rand"
	"strconv"
	"time"

	//"strconv"
)

// 转化为数字类型
func StringToInt(str string) int {
	if str == "" {
		log.Println("util.StringToInt 请勿输入空字符串")
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("util.StringToInt 字符串(%s)转换数字失败", str)
		return 0
	}

	return i
}

// 转化为正整数数字类型
func ToUint(str string) uint {
	i := StringToInt(str)

	return uint(i)
}

// 数字转化为字符串
func ToString(i int) string {
	return strconv.Itoa(i)
}

// interface 转为 string
func InterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	log.Println(key)
	return key
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

// 创建随机字符串
func CreateRandString(min int, max int, chars string) string {
	if chars == "" {
		chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_-"
	}

	length := 0
	if min == max && min == 0 {
		length = 100
	} else if min == max && min > 0 {
		length = min
	} else {
		rand.Seed(time.Now().Unix())
		length = rand.Intn(max - min) + min
	}

	str := ""
	char_len := len(chars)

	for length > 0 {
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := rd.Intn(char_len - 1)

		str += chars[i:i+1]

		length -= 1

		time.Sleep(1000 * 0.001)
	}

	return str
}