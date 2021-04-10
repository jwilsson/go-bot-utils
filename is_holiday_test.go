package utils

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	ascensionFriday := time.Date(2021, 05, 14, 12, 00, 00, 00, time.Local)
	result := IsHoliday(ascensionFriday)

	if result == false {
		t.Fatalf("Expected ascension Friday to be a holiday, but got '%t'", result)
	}

	nonHoliday := time.Date(2021, 02, 01, 12, 00, 00, 00, time.Local)
	result = IsHoliday(nonHoliday)

	if result != false {
		t.Fatalf("Expected non holiday to be 'false', but got '%t'", result)
	}
}
