package ascii

import (
	"testing"
	"fmt"
)

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}

func TestGreetingASCII(t *testing.T) {
	if isASCII(GreetingASCII()) == true {
		fmt.Println("Success: Inneholder bare ASCII chars")
		}else{
			t.Error("Feil: Inneholder extended chars(non ascii) ")
		}
	}
