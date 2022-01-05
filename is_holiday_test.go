package utils

import (
	"testing"
	"time"
)

func TestIsHolidayAscensionFriday(t *testing.T) {
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

func TestIsHolidayMonday(t *testing.T) {
	mondayBeforeRegularTuesday := time.Date(2021, 02, 01, 12, 00, 00, 00, time.Local)
	result := IsHoliday(mondayBeforeRegularTuesday)

	if result != false {
		t.Fatalf("Expected Monday before regular Tuesday not to be a holiday, but got '%t'", result)
	}

	mondayBeforeHolidayTuesday := time.Date(2023, 06, 06, 12, 00, 00, 00, time.Local)
	result = IsHoliday(mondayBeforeHolidayTuesday)
	if result == false {
		t.Fatalf("Expected Monday before holiday Tuesday to be a holiday, but got '%t'", result)
	}
}

func TestIsHolidayFriday(t *testing.T) {
	fridayAfterRegularThursday := time.Date(2022, 01, 14, 12, 00, 00, 00, time.Local)
	result := IsHoliday(fridayAfterRegularThursday)

	if result != false {
		t.Fatalf("Expected Friday after regular Thursday not to be a holiday, but got '%t'", result)
	}

	fridayAfterHolidayThursday := time.Date(2022, 01, 07, 12, 00, 00, 00, time.Local)
	result = IsHoliday(fridayAfterHolidayThursday)
	if result == false {
		t.Fatalf("Expected Friday after holiday Thursday to be a holiday, but got '%t'", result)
	}
}
