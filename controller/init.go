package controller

import "giligili/util"

func GetOffset(offset string) uint {
	i, err := util.ToInt(offset)
	if err != nil {
		panic(err)
	}

	return uint(i)
}

func GetLimit(limit string) uint {
	i, err := util.ToInt(limit)
	if err != nil {
		panic(err)
	}

	return uint(i)
}
