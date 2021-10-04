package code

import (
	"github.com/mattharr013/intcovpoc/strcheck"
	"testing"
)

func TestCodeDoesStuff_whenFalse(t *testing.T) {
	// @todo would we want this to count as coverage for the IsCool function?
	if DoesStuff(strcheck.IsCool("a string")) != "no" {
		t.Fail()
	}
}
