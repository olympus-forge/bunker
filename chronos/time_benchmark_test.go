package chronos

import (
	"testing"
	"time"
)

func BenchmarkAddWeekDay(b *testing.B) {
	today := time.Now()
	for i := 0; i < b.N; i++ {
		AddWeekDay(7, &today, nil)
	}
}
