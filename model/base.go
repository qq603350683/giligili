package model

import (
	"fmt"
	"time"
)

func IsDel(deltime time.Time) bool {
	if deltime != DelAtDefaultTime {
		return true
	}

	return false
}

func ClearTables(tables []string) bool {
	for _, table := range(tables) {
		DB.Exec("truncate " + table)

		fmt.Println("Clear table " + table)
	}

	return true
}