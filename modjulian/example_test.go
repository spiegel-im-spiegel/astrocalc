package modjulian_test

import (
	"fmt"
	"astrocalc/modjulian"
	"time"
)

func ExampleDayNumber() {
	t := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Print(modjulian.DayNumber(t))
	// Output:
	// 57023
}
