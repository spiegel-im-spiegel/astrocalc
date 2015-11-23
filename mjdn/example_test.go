package mjdn_test

import (
	"fmt"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/mjdn"
)

func ExampleDayNumber() {
	t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Print(mjdn.DayNumber(t))
	// Output:
	// 57023 (2015-01-01)
}

func ExampleToTime() {
	fmt.Print(mjdn.MJDN(57023).ToTime(time.UTC))
	// Output:
	// 2015-01-01 00:00:00 +0000 UTC
}
