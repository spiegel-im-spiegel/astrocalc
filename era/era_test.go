package era

import (
	"os"
	"testing"
	"time"

	"github.com/spiegel-im-spiegel/astrocalc/mjdn"
)

var (
	testTime = time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)
	testMJDN = mjdn.DayNumber(testTime)      //47534
	testEra  = New("Heisei", "平成", testMJDN) // 1989-01-08
)

func TestMain(m *testing.M) {
	//initialization

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}

func TestString(t *testing.T) {
	want := "平成 (Heisei) 1"
	result := testEra.String()
	if result != want {
		t.Errorf("String() = \"%s\", want \"%s\".", result, want)
	}
}

func TestCopy(t *testing.T) {
	result := testEra.Copy()
	if result.String() != testEra.String() {
		t.Errorf("Copy from \"%v\" to \"%v\", want \"%v\".", testEra, result, testEra)
	}
}

func TestGet(t *testing.T) {
	var testList = []struct {
		in  int64
		out time.Time
	}{
		{0, time.Date(1988, 1, 1, 0, 0, 0, 0, time.UTC)},
		{1, time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC)},
		{2, time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
		{27, time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, testCase := range testList {
		result := Get(testEra, testCase.in).ToTime()
		if result != testCase.out {
			t.Errorf("Era Year %d -> \"%v\", want \"%v\".", testCase.in, result, testCase.out)
		}
	}
}

func TestCheck(t *testing.T) {
	var testList = []struct {
		cmpr time.Time
		out  bool
	}{
		{time.Date(1989, 1, 7, 0, 0, 0, 0, time.UTC), false},
		{time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC), true},
		{time.Date(1989, 1, 9, 0, 0, 0, 0, time.UTC), true},
	}

	for _, testCase := range testList {
		result := testEra.Check(testCase.cmpr)
		if result != testCase.out {
			t.Errorf("Compare to \"%v\" = \"%v\", want \"%v\".", testCase.cmpr, result, testCase.out)
		}
	}
}

func TestToTime(t *testing.T) {
	result := testEra.ToTime()
	if result != testTime {
		t.Errorf("ToTime() = \"%v\", want \"%v\".", result, testTime)
	}
}

func TestToEra(t *testing.T) {
	var testList = []struct {
		in  time.Time
		out string
	}{
		{time.Date(1988, 12, 31, 0, 0, 0, 0, time.UTC), "平成 (Heisei) 0"},
		{time.Date(1989, 1, 7, 0, 0, 0, 0, time.UTC), "平成 (Heisei) 1"},
		{time.Date(1989, 1, 8, 0, 0, 0, 0, time.UTC), "平成 (Heisei) 1"},
		{time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), "平成 (Heisei) 2"},
		{time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), "平成 (Heisei) 27"},
	}

	for _, testCase := range testList {
		result := testEra.ToEra(testCase.in)
		if result.String() != testCase.out {
			t.Errorf("Time \"%v\" to Era \"%v\", want \"%v\".", testCase.in, result, testCase.out)
		}
	}
}
