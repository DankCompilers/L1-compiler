package L1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type codeGenerator interface {
	beginProcess(Node) error
}

type AsmCodeGenerator struct {
	writer *bufio.Writer
}

func newAsmCodeGenerator() *AsmCodeGenerator {
	return &AsmCodeGenerator{}
}

var calleeArray = [...]string{"%rbx", "%rbp", "%r12", "%r13", "%r14", "%r15"}
var callerArray = [...]string{"%rax", "%rcx", "%rdx", "%r8", "%r9", "%r10", "%r11"}

func (c *AsmCodeGenerator) compNode(node Node) {

	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ProgramNode:
		for _, reg := range calleeArray {
			c.writer.WriteString("push  %s", reg)
		}
		c.compNode(n.Child)

	case *subprog:

	}
}

//func (c *AsmCodeGenerator) WriteToFile(line string){
//}

func (c *AsmCodeGenerator) beginProcess(ast Node) error {
	file, err := os.Create("L1generatedASM.txt")
	defer file.Close()
	if err != nil {
		return err
	}

	c.writer = bufio.NewWriter(file)
	c.walkandTalk(ast)
	c.writer.Flush()
	return nil
}
