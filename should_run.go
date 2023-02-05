package utils

import (
	"strconv"
	"strings"
	"time"
)

func ShouldRun(currentTime time.Time, targetTime string, timezone string) bool {
	if IsHoliday(currentTime) {
		return false
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return false
	}

	localHour := currentTime.In(location).Hour()
	targetHour := strings.Split(targetTime, ":")[0]

	return targetHour == strconv.Itoa(localHour)
}
