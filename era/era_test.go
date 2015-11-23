package era

import (
	"os"
	"testing"
)

var (
	testEra = New("Heisei", "平成", -678577) // 0001-01-01
)

func TestMain(m *testing.M) {
	//initialization

	//start test
	code := m.Run()

	//termination
	os.Exit(code)
}
