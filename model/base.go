package model

import "time"

func IsDel(deltime time.Time) bool {
	if deltime != DelAtDefaultTime {
		return true
	}

	return false
}