package hash

import (
	"testing"
)

func TestDJBHash(t *testing.T) {
	testStr := "a"
	expect := uint(97)

	actual := DJBHash(testStr)
	if actual != expect {
		t.Error("djb hash error")
	}

	testStr = "abdedffdjasfjda"
	expect = 12856954179586838515
	actual = DJBHash(testStr)
	if actual != expect {
		t.Error("djb hash error")
	}
}
