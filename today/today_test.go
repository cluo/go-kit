package today

import (
	"testing"
	"time"
)

func TestTodayBeforeMidnight(t *testing.T) {
	year, month, day := time.Now().Date()

	v := TodayBeforeMidnight()

	// Test the date
	if year != v.Year() || month != v.Month() || day != v.Day() {
		t.Fatal("Date does not match: {Expected: %v-%v-%v, Actual: %v}", year, month, day, v)
	}

	// And the time
	if v.Hour() != 23 || v.Minute() != 59 || v.Second() != 59 || v.Nanosecond() != maxNano {
		t.Fatal("Invalid time returned: %v", v)
	}
}