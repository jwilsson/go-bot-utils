package utils

import (
	"testing"
	"time"
)

func TestShouldRun(t *testing.T) {
	location, _ := time.LoadLocation("Europe/Stockholm")

	date := time.Date(2022, 04, 04, 11, 15, 00, 00, location)
	result := ShouldRun(date, "11:15", location.String())

	if result != true {
		t.Fatalf("Expected same time result to be 'true', but got '%t'", result)
	}

	date = time.Date(2022, 04, 04, 12, 15, 00, 00, location)
	result = ShouldRun(date, "11:15", location.String())

	if result != false {
		t.Fatalf("Expected different times result to be 'false', but got '%t'", result)
	}
}
