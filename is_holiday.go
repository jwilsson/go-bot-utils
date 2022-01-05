package utils

import (
	"time"

	cal "github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/se"
)

func setup() *cal.BusinessCalendar {
	c := cal.NewBusinessCalendar()
	h := append(se.Holidays, &cal.Holiday{
		Name:   "Kristi himmelsfärdsdag (fredag)",
		Offset: 40,
		Func:   cal.CalcEasterOffset,
	})

	c.AddHoliday(h...)

	return c
}

func IsHoliday(t time.Time) bool {
	tomorrow := t.AddDate(0, 0, 1)
	yesterday := t.AddDate(0, 0, -1)
	weekday := t.Weekday().String()

	c := setup()

	if (weekday == "Monday" && !c.IsWorkday(tomorrow)) {
		return true
	}

	if (weekday == "Friday" && !c.IsWorkday(yesterday)) {
		return true
	}

	return !c.IsWorkday(t)
}
