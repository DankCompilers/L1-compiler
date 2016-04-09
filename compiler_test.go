package l1compiler

import (
	"testing"
)

func TestL1Compiler(t *testing.T) {
	c := NewCompiler()

	err := c.Compile("02.L1")
	if err != nil {
		t.Fatal(err)
	}

	err = c.GenerateCode()
	if err != nil {
		t.Fatal(err)
	}
}
