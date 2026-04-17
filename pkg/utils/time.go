package utils

import "time"

func GetEmptyTime() time.Time {
	return time.Time{}.AddDate(-1, -1, 0)
}
