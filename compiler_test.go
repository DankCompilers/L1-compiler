package l1compiler

import (
	"flag"
	"testing"
)

var fileToTest string

func init() {
	flag.StringVar(&fileToTest, "FileName", "none", "file name to be tested")
}

func TestL1Compiler(t *testing.T) {
	flag.Parse()
	c := NewCompiler()

	err := c.Compile(fileToTest)
	//	err := c.Compile("04.L1")
	if err != nil {
		t.Fatal(err)
	}

	err = c.GenerateCode()
	if err != nil {
		t.Fatal(err)
	}
}
