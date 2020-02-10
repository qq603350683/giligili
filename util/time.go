package util

import "time"

func Datetime(datetime time.Time) string {
	return datetime.Format("2006-01-02 15:04:05")
}