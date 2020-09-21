package controller

import "giligili/util"

func GetOffset(offset string) uint {
	i := util.StringToInt(offset)
	//if err != nil {
	//	panic(err)
	//}

	return uint(i)
}

func GetLimit(limit string) uint {
	i := util.StringToInt(limit)
	//if err != nil {
	//	panic(err)
	//}

	return uint(i)
}
