package modjulian

import (
	"os"
	"testing"
	"time"
)

type mjdnTest struct { //test case for DayNumber
	in  time.Time //input data
	out int64     //expected result
}

var mjdnTests []mjdnTest //test cases for DayNumber

func TestMain(m *testing.M) {
	//initialization
	mjdnTests = []mjdnTest{ //test cases for DayNumber
		{time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC), int64(40586)},
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), int64(40587)},
		{time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), int64(57023)},
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
