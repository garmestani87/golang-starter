package sum

import (
	"testing"
)

type TestCase struct {
	start    int
	end      int
	expexted int
}

func TestAdd(t *testing.T) {
	testCases := []TestCase{}
	testCases = append(testCases, TestCase{1, 10, 46})
	testCases = append(testCases, TestCase{1, 100, 4950})
	testCases = append(testCases, TestCase{1, 1000, 499500})

	for _, item := range testCases {
		actual := Add(item.start, item.end)
		if item.expexted != actual {
			t.Errorf("expected is %d but actual is %d", item.expexted, actual)
		}
	}
}

// go test
// go test -cover
