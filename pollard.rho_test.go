package primer

import (
	"testing"
)

func TestPollardRho(t *testing.T) {
	n := 10403
	result := PollardRho(n)
	if result != 101 {
		t.Error("Expected 101 got", result)
	}
}
