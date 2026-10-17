package chronos

import (
	"testing"
	"time"
)

func TestBunkerTime(t *testing.T) {
	t.Run("testAddWeekDay", func(t *testing.T) {
		expected := "2022-10-12"
		today, _ := time.Parse("2006-01-02", "2022-10-03")
		actualUnsetHoliday := AddWeekDay(7, &today, nil).Format("2006-01-02")
		if actualUnsetHoliday != expected {
			t.Errorf("invalid addWeekDay\n\tExpected : %v\n\tActual : %v", expected, actualUnsetHoliday)
		}

		// Holiday on a weekend (2022-10-08 is Saturday) - should not double-skip
		weekendHoliday, _ := time.Parse("2006-01-02", "2022-10-08")
		actualWeekendHoliday := AddWeekDay(7, &today, []time.Time{
			weekendHoliday,
		}).Format("2006-01-02")
		if actualWeekendHoliday != expected {
			t.Errorf("invalid addWeekDay with weekend holiday\n\tExpected : %v\n\tActual : %v", expected, actualWeekendHoliday)
		}

		// Holiday on a weekday (2022-10-05 is Wednesday) - should push SLA to 2022-10-13
		weekdayHoliday, _ := time.Parse("2006-01-02", "2022-10-05")
		expectedWithHoliday := "2022-10-13"
		actualWeekdayHoliday := AddWeekDay(7, &today, []time.Time{
			weekdayHoliday,
		}).Format("2006-01-02")
		if actualWeekdayHoliday != expectedWithHoliday {
			t.Errorf("invalid addWeekDay with weekday holiday\n\tExpected : %v\n\tActual : %v", expectedWithHoliday, actualWeekdayHoliday)
		}

		// Zero SLA should return today, not zero time 0001-01-01
		actualZeroSla := AddWeekDay(0, &today, nil).Format("2006-01-02")
		if actualZeroSla != "2022-10-03" {
			t.Errorf("invalid addWeekDay for 0 sla\n\tExpected : 2022-10-03\n\tActual : %v", actualZeroSla)
		}

		// Test fluent New() chaining
		fluentDate := New().SetToday(&today).SetHolidays([]time.Time{weekdayHoliday}).AddWeekDay(7).Format("2006-01-02")
		if fluentDate != expectedWithHoliday {
			t.Errorf("invalid fluent AddWeekDay\n\tExpected : %v\n\tActual : %v", expectedWithHoliday, fluentDate)
		}
	})
}
