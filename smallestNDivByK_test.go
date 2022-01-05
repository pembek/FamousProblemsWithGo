package FamousProblemsWithGo
// go test smallestNDivByK.go smallestNDivByK_test.go
import (
	"testing"
)

func TestSmallestNDivByK(t *testing.T) {
	res := SmallestNDivByK(3)

	if res != 3 {
		t.Errorf("Expected 3 for 111")
	}

	res = SmallestNDivByK(1)

	if res != 1 {
		t.Errorf("Expected 1 for 1")
	}

	res = SmallestNDivByK(2)

	if res != -1 {
		t.Errorf("Expected -1")
	}

	res = SmallestNDivByK(23)

	if res != 22 {
		t.Errorf("Expected 22 for 23")
	}

	// pigeon hole principle
	res = SmallestNDivByK(19927)

	if res != 19926 {
		t.Errorf("Expected 19926 for 19927")
	}

	res = SmallestNDivByK(5367)

	if res != 1788 {
		t.Errorf("Expected 1788 for 5367")
	}

	res = SmallestNDivByK(1983)

	if res != 660 {
		t.Errorf("Expected 660 for 1983")
	}
}
