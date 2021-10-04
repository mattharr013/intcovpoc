package integration

import (
	"github.com/mattharr013/intcovpoc/code"
	"github.com/mattharr013/intcovpoc/strcheck"
	"testing"
)

func TestCodeDoesStuff_whenTrue(t *testing.T) {
	if code.DoesStuff(strcheck.IsEmpty("")) != "yes" {
		t.Fail()
	}
}
