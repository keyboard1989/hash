package hash

import (
	"testing"
)

func TestDJBHash(t *testing.T) {
	testStr := "a"
	expect := uint(177670)

	actual := DJBHash(testStr)
	if actual != expect {
		t.Error("djb hash error")
	}

	testStr = "abdedffdjasfjda"
	expect = 9018872858425983576
	actual = DJBHash(testStr)
	if actual != expect {
		t.Error("djb hash error")
	}
}
