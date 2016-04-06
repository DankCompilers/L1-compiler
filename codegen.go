package l1compiler

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var calleeArray = [...]string{"%rbx", "%rbp", "%r12", "%r13", "%r14", "%r15"}
var callerArray = [...]string{"%rax", "%rcx", "%rdx", "%r8", "%r9", "%r10", "%r11"}

var eightBitEquiv = map[string]string{
	"r10": "r10b",
	"r13": "r13b",
	"r8":  "r8b",
	"rbp": "bpl",
	"rdi": "dil",
	"r11": "r11b",
	"r14": "r14b",
	"r9":  "r9b",
	"rbx": "bl",
	"rdx": "dl",
	"r12": "r12b",
	"r15": "r15b",
	"rax": "al",
	"rcx": "cl",
	"rsi": "sil",
}

func labelToASM(label string) string {

	newLabel := strings.TrimLeft(label, ":")
	newLabel = "_" + newLabel

	return newLabel
}

type CodeGenerator interface {
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
		c.writer.WriteString("go: \n")
		for _, reg := range calleeArray {
			toWrite := fmt.Sprintf("push  %%s \n", reg)
			c.writer.WriteString(toWrite)
		}

		c.compNode(n.Front())

		for i := len(calleeArray) - 1; i >= 0; i-- {
			regToWrite := calleeArray[i]
			toWrite := fmt.Sprintf("pop  %%s \n", regToWrite)
			c.writer.WriteString(toWrite)
		}
		c.writer.WriteString("retq \n")

		//	case *AssignNode:
		//		mem, value := n.children[:]
		//		c.writer.WriteString("movq $%d, %s \n", value, mem)
	case nil:
		//return
	case *SubProgramNode, *InstructionNode:
		dummyNode := n.Front()
		c.compNode(dummyNode)
		for {
			c.compNode(n.Next())
		}

	case *FunctionNode:
		label := n.Label
		locals := n.NumLocals
		label = labelToASM(label)
		label = label + "\n"
		c.writer.WriteString(label)
		if locals > 0 {
			toWrite := fmt.Sprintf("subq $%d, %rsp", (8 * locals))
			c.writer.WriteString(toWrite)
		}
		dummyNode := n.Front()
		c.compNode(dummyNode)
		for {
			c.compNode(n.Next())
		}

	case *OpNode:
		mem := n.Front().returnLabel()
		value := n.Next().returnLabel()

		var instruct string

		if op := n.Operator; op == ">>=" {
			instruct = "sarq"
		} else {
			instruct = "salq"
		}

		if numRep, err := strconv.Atoi(value); err == nil {
			value = "$" + strconv.Itoa(numRep)
		}

		lowerReg := eightBitEquiv[mem]
		toWrite := fmt.Sprintf("%s %s, %s", instruct, value, lowerReg)
		c.writer.WriteString(toWrite)

		//	case *CallNode:

		//		u, NAT = n.children[:]
		//		u := u.value
		//		NAT := NAT

	case *LabelNode:
		newString := labelToASM(n.Label)
		//return newString
		//c.writer.WriteString(newString + "\n")

	case *GotoNode, *SysCallNode:
		newString := labelToASM(n.returnLabel())
		newString = "jmp " + newString
		c.writer.WriteString(newString)

	case *CallNode:
		if arity := n.Arity; arity > 6 {
			spills := 8 * (arity - 6)
			toWrite := fmt.Sprintf("subq $%d %s", spills, "%rsp")
			c.writer.WriteString(toWrite)
		}
		returnedString := c.compNode(n.Front())
		toWrite := fmt.Sprintf("jmp %s", returnedString)
		c.writer.WriteString(toWrite)

	case *AssignNode:
		if child := n.Front(); reflect.TypeOf(child) == MemNode {
			reg := child.X
			offset := child.N8
			throwAway := n.Front()
			leftChildValue := n.Next().returnLabel()
			toInput := "%" + leftChildValue
			if i, err := strconv.Atoi(leftChildValue); err == nil {
				toInput = "$" + strconv.Itoa(i)
			}

			toWrite := fmt.Sprintf("movq %s, %d(%s) \n", toInput, offset, reg)
			c.writer.WriteString(toWrite)
		} else {

			destination := n.Front().returnLabel()
			if lChild := n.Next(); lChild.(type) != *MemNode {

				inputToStore := c.compNode(lChild)
				tokenVal := lChild.returnLabel()
				if _, ok := eightBitEquiv[tokenVal]; ok {
					//  c.writer.Writer("movq %%s, %%s", tokenVal, destination)
					inputToStore = "%" + tokenVal

				} else if i, err := strconv.Atoi(tokenVal); err == nil {

					inputToStore = "$" + tokenVal
				}

				toWrite := fmt.Sprintf("movq %s, %s \n", inputToStore, destination)
				c.writer.WriteString(toWrite)
			} else if lChild.(type) == *CmpopNode {
				onleft, val := c.compNode(lChild)
				rChild := n.Front().returnLabel()
				if onleft == 0 {
					toWrite := fmt.Sprintf("movq $%s, %%s", rChild)
					c.writer.WriteString(toWrite)
				} else {
					eightMapped := eightBitEquiv[rChild]
					toWrite := fmt.Sprintf("%s %s", val, eightMapped)
					c.writer.WriteString(toWrite)
					toWrite = fmt.Sprintf("movzbq %s, %s", val, rChild)
					c.writer.WriteString(toWrite)
				}

			} else {
				reset := n.Front()
				memoryNode := n.Next()
				memX := memoryNode.X
				memOffset := memoryNode.N8

				toWrite := fmt.Sprintf("movq %d(%%s), %s \n", memOffset, memX, destination)
				c.writer.WriteString(toWrite)
			}

		}

	case *CmpopNode:
		op := Operator
		leftSide := n.Front().token
		rightSide = n.Next().token

		switch op {
		default:
			opCode := nil
		case "<":
			opCode := "setl"
		case "<=":
			opCode := "setge"
		case ">=":
			opCode := "setle"
		case "=":
			opCode := "sete"
		}

		if ls, ok := strconv.Atoi(leftSide); ok {
			if rs, ok := strconv.Atoi(rightSide); ok {
				switch op {

				default:
					return

				case "<":
					boolToReturn := (ls < rs)
					return 0, boolToReturn, op

				case "<=":
					boolToReturn := (ls <= rs)
					return 0, boolToReturn, op

				case "=":
					boolToReturn := (ls == rs)
					return 0, boolToReturn, op
				}
			} else {
				toWrite = fmt.Sprintf("cmpq $%d, %%s \n", ls, rightSide)
				c.writer.WriteString(toWrite)
				switch opCode {
				default:
					nil
				case "setl":
					opCode = "setg"
				case "setg":
					opCode = "setl"
				case "setge":
					opCode = "setle"
				case "setle":
					opCode = "setge"
				}
				return 1, opCode, op
			}
		} else {
			if _, ok = strconv.Atoi(leftSide); ok {
				leftSide = "$" + leftSide
			}

			toWrite = fmt.Sprintf("cmpq %s, %%s \n", leftSide, rightSide)
			c.writer.WriteString(toWrite)
			return 1, opCode, op
		}

	case *CjumpNode:
		onleft, val, op = c.compNode(n.Front)
		yesLabel := n.Next().label
		noLabel := n.Next().label
		switch op {
		case "<=":
			jmpStment = "jle"
		case "<":
			jmpStment = "jl"
		case "=":
			jmpStment = "je"
		case ">=":
			jmpStment = "jge"
		}
		toWrite = fmt.Sprintf("%s %s", jmpStment, yesLabel)
		c.writer.WriteString(toWrite)
		toWrite = fmt.Sprintf("jmp %s", noLabel)
		c.writer.WriteString(toWrite)

	}
	return
}

func (c *AsmCodeGenerator) BeginCompiler(ast Node) error {
	file, err := os.Create("L1generatedASM.txt")
	defer file.Close()
	if err != nil {
		return err
	}

	c.writer = bufio.NewWriter(file)
	c.compNode(ast)
	c.writer.Flush()
	return nil
}
