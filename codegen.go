package l1compiler

import (
	"bufio"
	"fmt"
	"os"
	//	"reflect"
	"strconv"
	"strings"
)

var calleeArray = [...]string{"%rbx", "%rbp", "%r12", "%r13", "%r14", "%r15", "%rdi"}
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

var aopMAP = map[string]string{
	"+=": "addq",
	"-=": "subq",
	"*=": "imulq",
	"&=": "andq",
}

var sopMap = map[string]string{
	">>=": "sarq",
	"<<=": "salq",
}

func labelToASM(label string) string {
	if _, err := strconv.Atoi(label); err == nil {
		inputToStore := "$" + label
		return inputToStore

	} else if _, ok := eightBitEquiv[label]; ok {
		//c.writer.WriteString("movq %%s, %%s", tokenVal, destination)
		//c.writer.WriteString("in reg case \n")
		inputToStore := "%" + label
		return inputToStore
	} else {
		newLabel := strings.TrimLeft(label, ":")
		newLabelNew := "_" + newLabel + ":"

		return newLabelNew
	}
}

type CodeGenerator interface {
	BeginCompiler(Node) error
}

type AsmCodeGenerator struct {
	writer *bufio.Writer
}

func L1toASMGenerator() *AsmCodeGenerator {
	return &AsmCodeGenerator{}
}

var returnHolders []interface{}
var defaultRet []interface{} = []interface{}{0, "nil", "nil"}

func (c *AsmCodeGenerator) compNode(node Node) (int, string, string) {

	if node == nil {
		//def := &defaultRet
		return 0, "0", "0"
	}

	switch n := node.(type) {

	case *ProgramNode:
		c.writer.WriteString(".text\n")
		//preMain = n.Label
		main := n.Label
		toWrite := fmt.Sprintf(".globl %s\n", main)
		c.writer.WriteString(toWrite)
		toWrite = fmt.Sprintf("%s:\n", main)
		c.writer.WriteString(toWrite)
		for _, reg := range calleeArray {
			toWrite := fmt.Sprintf("push  %s\n", reg)
			c.writer.WriteString(toWrite)
		}

		c.compNode(n.Front())

		for i := len(calleeArray) - 1; i >= 0; i-- {
			regToWrite := calleeArray[i]
			toWrite := fmt.Sprintf("pop  %s\n", regToWrite)
			c.writer.WriteString(toWrite)
		}
		c.writer.WriteString("retq\n")

		//	case *AssignNode:
		//		mem, value := n.children[:]
		//		c.writer.WriteString("movq $%d, %s \n", value, mem)
	case nil:
		//return
	case *SubProgramNode:
		dummyNode := n.Front()
		c.compNode(dummyNode)
		iter := len(n.children)
		for i := 0; i < iter-1; i++ {
			c.compNode(n.Next())
		}

	case *InstructionNode:
		dummyNode := n.Front()
		c.compNode(dummyNode)
		iter := len(n.children)
		for i := 0; i < iter-1; i++ {
			c.compNode(n.Next())
		}

	case *FunctionNode:
		label := n.Label
		locals := n.NumLocals
		label = labelToASM(label)
		label = label + "\n"
		c.writer.WriteString(label)
		if locals > 0 {
			toWrite := fmt.Sprintf("subq $%d, %rsp \n", (8 * locals))
			c.writer.WriteString(toWrite)
		}
		dummyNode := n.Front()
		c.compNode(dummyNode)
		iter := len(n.children)

		for i := 0; i < iter-1; i++ {
			c.compNode(n.Next())
		}

	case *OpNode:
		_, _, op1 := c.compNode(n.Front())
		//value := n.Next().returnLabel()
		_, _, op2 := c.compNode(n.Next())

		var shiftAmount string
		var toWrite string
		/*
			if op := n.Operator; op == ">>=" {
				instruct = "sarq"
			} else {
				instruct = "salq"
			}
			toWrite := fmt.Sprintf("%s %s, %s\n", instruct, value, mem)
			c.writer.WriteString(toWrite)
		*/
		if aop, ok := aopMAP[n.Operator]; ok {
			toWrite = fmt.Sprintf("%s, %s, %s\n", aop, op2, op1)
		} else {
			sop := sopMap[n.Operator]
			label := n.children[n.currChild].returnLabel()
			if shiftBy, ok := eightBitEquiv[label]; ok {
				shiftAmount = shiftBy
			} else {
				shiftAmount = op2
			}
			toWrite = fmt.Sprintf("%s, %s, %s\n", sop, shiftAmount, op1)
		}

		c.writer.WriteString(toWrite)

	case *ReturnNode:
		c.writer.WriteString("ret\n")

	case *LabelNode:
		newString := labelToASM(n.Label)
		return 0, "0", newString
		//c.writer.WriteString(newString + "\n")

	case *GotoNode:
		newString := labelToASM(n.Label)
		newString = fmt.Sprintf("jmp %s\n", newString)
		c.writer.WriteString(newString)
		return 0, "nil", "nil"

	case *SysCallNode:
		//newString := labelToASM(n.Label)
		newStringFinal := fmt.Sprintf("call %s\n", n.Label)
		c.writer.WriteString(newStringFinal)
		return 0, "nil", "nil"

	case *CallNode:
		if arity := n.Arity; arity > 6 {
			spills := 8 * (arity - 6)
			toWrite := fmt.Sprintf("subq $%d %s\n", spills, "%rsp")
			c.writer.WriteString(toWrite)
		}
		_, _, returnedString := c.compNode(n.Front())
		toWrite := fmt.Sprintf("jmp %s\n", returnedString)
		c.writer.WriteString(toWrite)
		return 0, "nil", "nil"

	case *AssignNode:
		switch child := n.Front().(type) {
		case *MemNode:
			//	if child := n.Front(); reflect.TypeOf(child) == MemNode {
			reg := child.X
			offset := child.N8
			_ = n.Front()
			leftChildValue := n.Next().returnLabel()
			toInput := "%" + leftChildValue
			if i, err := strconv.Atoi(leftChildValue); err == nil {
				toInput = "$" + strconv.Itoa(i)
			}

			toWrite := fmt.Sprintf("movq %s, %d(%s)\n", toInput, offset, reg)
			c.writer.WriteString(toWrite)
			return 0, "nil", "nil"

		default:

			_, _, destination := c.compNode(n.Front())
			switch lChild := n.Next().(type) {

			default:
				_, _, inputToStore := c.compNode(lChild)
				/*
					tokenVal := lChild.returnLabel()
					//tokenVal := "rdi"
					//fmt.Printf(tokenVal)
					if _, ok := eightBitEquiv[tokenVal]; ok {
						//c.writer.WriteString("movq %%s, %%s", tokenVal, destination)
						c.writer.WriteString("in reg case \n")
						//inputToStore = "%" + tokenVal

					} else if _, err := strconv.Atoi(tokenVal); err == nil {
						c.writer.WriteString("in num case \n")
						inputToStore = "$" + tokenVal
					}
					//c.writer.WriteString(destination)
					//toSing := fmt.Sprintf("here is tokenVal %s \n", tokenVal)
					//c.writer.WriteString(toSing)
				*/
				toWrite := fmt.Sprintf("movq %s, %s\n", inputToStore, destination)
				c.writer.WriteString(toWrite)

				return 0, "nil", "nil"

			case *CmpopNode:
				onleft, val, _ := c.compNode(lChild) //*****
				rChild := n.Front().returnLabel()
				if onleft == 0 {
					toWrite := fmt.Sprintf("movq $%s, %%s\n", rChild)
					c.writer.WriteString(toWrite)
				} else {
					eightMapped := eightBitEquiv[rChild]
					toWrite := fmt.Sprintf("%s %s\n", val, eightMapped)
					c.writer.WriteString(toWrite)
					toWrite = fmt.Sprintf("movzbq %s, %s\n", val, rChild)
					c.writer.WriteString(toWrite)
				}

				return 0, "nil", "nil"

			case *MemNode:

				//reset := n.Front()
				memoryNode := lChild.Front()
				memX := memoryNode.returnLabel()
				memOffset := lChild.N8

				toWrite := fmt.Sprintf("movq %d(%%s), %s\n", memOffset, memX, destination)
				c.writer.WriteString(toWrite)

				return 0, "nil", "nil"
			}

		}

	case *CmpopNode:
		op := n.returnCmpop()
		leftSide := n.Front().returnLabel()
		rightSide := n.Next().returnLabel()

		opCode := "nil"

		switch op {
		//		default:
		//			opCode := nil
		case "<":
			opCode = "setl"
		case "<=":
			opCode = "setge"
		case ">=":
			opCode = "setle"
		case "=":
			opCode = "sete"
		}

		if ls, ok := strconv.Atoi(leftSide); ok == nil {
			if rs, ok := strconv.Atoi(rightSide); ok == nil {
				switch op {

				default:
					return 0, "nil", "nil"

				case "<":
					boolToReturn := (ls < rs)
					boolT := strconv.FormatBool(boolToReturn)
					return 0, boolT, op

				case "<=":
					boolToReturn := (ls <= rs)
					boolT := strconv.FormatBool(boolToReturn)
					return 0, boolT, op

				case "=":
					boolToReturn := (ls == rs)
					boolT := strconv.FormatBool(boolToReturn)
					return 0, boolT, op
				}
			} else {
				toWrite := fmt.Sprintf("cmpq $%d, %%s\n", ls, rightSide)
				c.writer.WriteString(toWrite)
				switch opCode {
				//				default:
				//					nil
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
			if _, ok = strconv.Atoi(leftSide); ok == nil {
				leftSide = "$" + leftSide
			}

			toWrite := fmt.Sprintf("cmpq %s, %%s\n", leftSide, rightSide)
			c.writer.WriteString(toWrite)
			return 1, opCode, op
		}

	case *CjumpNode:
		_, _, op := c.compNode(n.Front())
		yesLabel := n.Next().returnLabel()
		noLabel := n.Next().returnLabel()

		var jmpStment string
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

		toWrite := fmt.Sprintf("%s %s\n", jmpStment, yesLabel)
		c.writer.WriteString(toWrite)
		toWrite = fmt.Sprintf("jmp %s\n", noLabel)
		c.writer.WriteString(toWrite)

	}
	return 0, "nil", "nil"
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
