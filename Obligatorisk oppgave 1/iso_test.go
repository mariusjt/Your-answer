package iso

import (
	"testing"
)

func TestGretingsExtendedASCII(t *testing.T) {
	ascii := GreetingExtendedASCII()
	if !(IsASCII(ascii)) {
		t.Fail()
	}
}

func IsASCII(s string) bool {
	for _, C := range s {
		if C > 127 && C < 256 {
			return true
		}
	}
	return false
}
