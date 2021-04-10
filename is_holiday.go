package utils

import (
	"time"

	cal "github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/se"
)

func IsHoliday(t time.Time) bool {
	c := cal.NewBusinessCalendar()
	h := append(se.Holidays, &cal.Holiday{
		Name:   "Kristi himmelsfärdsdag (fredag)",
		Offset: 40,
		Func:   cal.CalcEasterOffset,
	})

	c.AddHoliday(h...)

	return !c.IsWorkday(t)
}
