package L1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var calleeArray = [...]string{"%rbx", "%rbp", "%r12", "%r13", "%r14", "%r15"}
var callerArray = [...]string{"%rax", "%rcx", "%rdx", "%r8", "%r9", "%r10", "%r11"}

eightBitEquiv := map[string]string{
"r10" : "r10b",
"r13" : "r13b",
"r8"  : "r8b",
"rbp" : "bpl",
"rdi" : "dil",
"r11" : "r11b",
"r14" : "r14b",
"r9"  : "r9b",
"rbx" : "bl",
"rdx" : "dl",
"r12" : "r12b",
"r15" : "r15b",
"rax" : "al",
"rcx" : "cl",
"rsi" : "sil",
}

func labelToASM(label string) string {

		newLabel := strings.TrimLeft(label, ':')
		newLabel := '_' + newString
}

type codeGenerator interface {
	beginCompiler(Node) error
}

type AsmCodeGenerator struct {
	writer *bufio.Writer
}

func L1toASMGenerator() *AsmCodeGenerator {
	return &AsmCodeGenerator{}
}

func (c *AsmCodeGenerator) compNode(node Node) {

	if node == nil {
		return
	}

	switch n := node.(type) {

	case *ProgramNode:
		c.writer.WriteString(".text \n")
		c.writer.WriteString(".globl go \n")
		c.writer.Writestring("go: \n")
		for _, reg := range calleeArray {
			c.writer.WriteString("push  %%s \n", reg)
		}

		c.compNode(n.children[0])

		for reg :=len(calleeArray) - 1; i>=0; i-- {
			regToWrite := calleeArray[i]
			c.writer.WriteString("pop  %%s \n", regToWrite)
		}
		c.writer.WriteString("retq \n")

	case *AssignNode:
		mem, value := n.children[:]
		c.writer.WriteString("movq $%d, %s \n", value, mem)

	case *OpNode:
		mem, value := n.children[:]

		if self := n.self; self == ">>=" {
			instruct := "sarq"
		} else {
			instruct := "salq"
		}

		if  numRep,err := strconv.Atoi(value); err == nil {
			value = '$' + numRep
		}

		lowerReg = eightBitEquiv[mem]
		c.writer.WriteString("%s %s, %s", instruct, value, lowerReg)

	case *CallNode:

		u,NAT = n.children[:]
		u := u.value
		NAT := NAT

	case *LabelNode:

		newString := labelToASM(n.label)
		c.writer.WriteString(newString + "\n")

	case *GotoNode:
		newString := labelToASM(n.label)
		c.writer.WriteString("jmp %s", newString)

	}
}

//func (c *AsmCodeGenerator) WriteToFile(line string){
//}

func (c *AsmCodeGenerator) beginCompiler(ast Node) error {
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
