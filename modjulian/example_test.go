package modjulian_test

import (
	"fmt"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/modjulian"
)

func ExampleDayNumber() {
	t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Print(modjulian.DayNumber(t))
	// Output:
	// 57023
}

func ExampleToTime() {
	fmt.Print(modjulian.ToTime(int64(57023), time.UTC))
	// Output:
	// 2015-01-01 00:00:00 +0000 UTC
}
