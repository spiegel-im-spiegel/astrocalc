package mjdn

import (
	"os"
	"testing"
	"time"
)

type mjdnTest struct { //test case for DayNumber
	in  time.Time //input data
	out MJDN      //expected result
}

var mjdnTests []mjdnTest //test cases for DayNumber

func TestMain(m *testing.M) {
	//initialization
	mjdnTests = []mjdnTest{ //test cases for DayNumber
		{time.Date(-4712, 1, 2, 0, 0, 0, 0, time.UTC), MJDN(-2400000)},
		{time.Date(-1, 1, 1, 0, 0, 0, 0, time.UTC), MJDN(-679308)},
		{time.Date(0, 10, 31, 0, 0, 0, 0, time.UTC), MJDN(-678639)},
		{time.Date(0, 12, 31, 0, 0, 0, 0, time.UTC), MJDN(-678578)},
		{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), MJDN(-678577)},
		{time.Date(1582, 10, 4, 0, 0, 0, 0, time.UTC), MJDN(-100841)},
		{time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC), MJDN(-100840)},
		{time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC), MJDN(40586)},
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), MJDN(40587)},
		{time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC), MJDN(47534)}, //平成元年
		{time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), MJDN(57023)},
	}

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestDayNumber(t *testing.T) {
	for _, testCase := range mjdnTests {
		result := DayNumber(testCase.in)
		if result != testCase.out {
			t.Errorf("DayNumber of \"%v\" = %d, want %d.", testCase.in, result, testCase.out)
		}
	}
}

func TestToTime(t *testing.T) {
	for _, testCase := range mjdnTests {
		result := testCase.out.ToTime(testCase.in.Location())
		if result != testCase.in {
			t.Errorf("time.Time of %d = \"%v\", want \"%v\".", testCase.out, result, testCase.in)
		}
	}
}
