package utils

import (
	"fmt"
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
	targetHour, err := strconv.Atoi(strings.Split(targetTime, ":")[0])
	if err != nil {
		return false
	}

	return fmt.Sprintf("%01d", targetHour) == strconv.Itoa(localHour)
}
