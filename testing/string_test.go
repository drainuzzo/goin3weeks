package strings_test

import (
	"fmt"
	"stringutils"
	"testing"
)

func TestUpper(t *testing.T) {
	word := "ciccio"
	if word.Upper != word.ToUpper {

		t.Errorf("Word not Upperata")
	}
}

func TestLower(t *testing.T) {
	word := "ciccio"
	if word.Lower != word.ToLower {

		t.Errorf("Word not Lowerata")
	}
}
