package l1compiler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var calleeArray = map[string]bool{
   "%rbx": false,
   "%rbp": false,
   "%r12": false,
   "%r13": false,
   "%r14": false,
   "%r15": false,
}

var callerArray = map[string]bool{
   "%rax": false,
   "%rcx": false,
   "%rdx": false,
   "%rsi": false,
   "%rdi": false,
   "%r8": false,
   "%r9": false,
   "%r10": false,
   "%r11": false,
}

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

var allRegisters = map[string]bool{
	"r10": true,
	"r13": true,
	"r8":  true,
	"rbp": true,
	"rdi": true,
	"r11": true,
	"r14": true,
	"r9":  true,
	"rbx": true,
	"rdx": true,
	"r12": true,
	"r15": true,
	"rax": true,
	"rcx": true,
	"rsi": true,
	"rsp": true,
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



var functionList = map[string]*FunctionNode{}
var labelList = map[string][]string{}
var functionCalls = map[string]*CallNode{}


func labelToASM(label string) string {
	if _, err := strconv.Atoi(label); err == nil {
		inputToStore := "$" + label
		return inputToStore

	} else if _, ok := allRegisters[label]; ok {
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

func saveCalleeRegisters() string {
   toWrite := []string{}

   for _, reg := range calleeArray {
	  toWrite = append(toWrite, fmt.Sprintf("push  %s\n", reg))
   }

   return strings.Join(toWrite)
}

func (c *AsmCodeGenerator) compNode(node Node) (int, string, string) {

	if node == nil {
		return 0, "0", "0"
	}

	switch n := node.(type) {

	case *ProgramNode:
		c.writer.WriteString(".text\n")
		main := fmt.Sprintf("_%s", n.Label[1:]) //removes colon
		toWrite := fmt.Sprintf(".globl %s\n\n%s:\n", main,main)
		c.writer.WriteString(toWrite)

		 c.writer.WriteString(saveCalleeRegisters())

		c.writer.WriteString(fmt.Sprintf("call %s\n\n", main))

		c.compNode(n.Front())
		c.writer.WriteString("\n\n")

		for i := len(calleeArray) - 1; i >= 0; i-- {
			regToWrite := calleeArray[i]
			toWrite := fmt.Sprintf("pop  %s\n", regToWrite)
			c.writer.WriteString(toWrite)
		}

		c.writer.WriteString("retq\n")

	case *SubProgramNode:
		dummyNode := n.Front()
		c.compNode(dummyNode)
		iter := len(n.children)

		for i := 0; i < iter-1; i++ {
			c.compNode(n.Next())
			c.writer.WriteString("\n")
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
		functionList[label] = n

		label = label
		c.writer.WriteString(label + "\n")

	    spills := 0
		 if arity := n.Arity; arity > 6 {
			spills += 8 * (int(arity) - 6)
		 }

		 if locals > 0 {
			spills += int(locals) * 8
		 }

		 if spills > 0 {
			toWrite := fmt.Sprintf("subq $%d, %%rsp \n", (8 * locals))
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
		_, _, op2 := c.compNode(n.Next())

		var shiftAmount string
		var toWrite string

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


		_, _, returnedString := c.compNode(n.Front())
		toWrite := fmt.Sprintf("jmp %s\n", returnedString)
		toWriter = fmt.Sprintf("%s\n"
		c.writer.WriteString(toWrite)
		return 0, "nil", "nil"

   case *TokenNode:
	  return 0, "0", labelToASM(n.Value)
	case *AssignNode:
		switch child := n.Front().(type) {

		case *MemNode:

			_, _, reg := c.compNode(child.Front())
			offset := child.N8
			_ = n.Front()
			_, _, toInput := c.compNode(n.Next())
			toWrite := fmt.Sprintf("movq %s, %d(%s)\n", toInput, offset, reg)
			c.writer.WriteString(toWrite)
			return 0, "nil", "nil"

		default:

			_, _, destination := c.compNode(n.Front())
			switch lChild := n.Next().(type) {

			default:
				_, _, inputToStore := c.compNode(lChild)
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
