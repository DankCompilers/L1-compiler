package l1compiler

import (
	"bufio"
	"fmt"
	"os"
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

var return_address string
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

type spillManage struct {
	spillArray []int
	Amount     int
}

type spillBank interface {
	spill_accrue(int)
	spill_thief() int
}

func (spills *spillManage) spill_accrue(spill int) {
	spills.spillArray = append(spills.spillArray, spill)
	spills.Amount += 1
}

func (spills *spillManage) spill_thief() int {
	val2RET := spills.spillArray[spills.Amount]
	//spills.Amount -= 1
	spills.spillArray = spills.spillArray[0:spills.Amount]
	spills.Amount -= 1
	return val2RET
}

var spill_manage = new(spillManage)

var spill_bank = spillBank(spill_manage)

func saveCalleeRegisters() string {
	toWrite := []string{}
	for _, reg := range calleeArray {
		toWrite = append(toWrite, fmt.Sprintf("push  %s\n", reg))
	}

	toWriteFinal := strings.Join(toWrite, "")
	return toWriteFinal
}

func popCalleeRegisters() string {
	size := len(calleeArray)
	toWrite := make([]string, size)
	var i = (size - 1)
	for _, reg := range calleeArray {
		registerPop := fmt.Sprintf("pop %s\n", reg)
		toWrite[i] = registerPop
		i--

	}

	toWriteFinal := strings.Join(toWrite, "")
	return toWriteFinal
}

func (c *AsmCodeGenerator) compNode(node Node) (int, string, string) {

	if node == nil {
		return 0, "0", "0"
	}

	switch n := node.(type) {

	case *ProgramNode:
		c.writer.WriteString(".text\n")
		main := fmt.Sprintf("%s", n.Label[1:]) //removes colon
		//toWrite := fmt.Sprintf(".globl %s\n\n%s:\n", main, main)
		toWrite := fmt.Sprintf(".globl %s\n\n%s\n", "go", "go:")
		c.writer.WriteString(toWrite)

		c.writer.WriteString(saveCalleeRegisters())
		//main = main + ":"
		c.writer.WriteString(fmt.Sprintf("call _%s\n\n", main))
		c.writer.WriteString("\n\n")
		c.writer.WriteString(popCalleeRegisters())
		c.writer.WriteString("retq\n")

		c.compNode(n.Front())
		c.writer.WriteString("\n\n")

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
			spills += (8 * (int(arity) - 6))
		}

		if locals > 0 {
			spills += int(locals) * 8
			toWrite := fmt.Sprintf("subq $%d, %%rsp \n", (8 * locals))
			c.writer.WriteString(toWrite)
		}

		if spills > 0 {
			spill_manage.spill_accrue(spills)
			c.writer.WriteString("###CALLING SPILL ACCRUE###\n")

		} else {
			spill_manage.spill_accrue(0)
			c.writer.WriteString("###CALLING SPILL ACCRUE###\n")
		}
		c.writer.WriteString("###IN FUNCTION NODE###\n")

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
			toWrite = fmt.Sprintf("%s %s, %s\n", aop, op2, op1)
		} else {
			sop := sopMap[n.Operator]
			label := n.children[n.currChild].returnLabel()
			if shiftBy, ok := eightBitEquiv[label]; ok {
				shiftAmount = shiftBy
			} else {
				shiftAmount = op2
			}
			toWrite = fmt.Sprintf("%s %s, %s\n", sop, shiftAmount, op1)
		}

		c.writer.WriteString(toWrite)

	case *ReturnNode:
		var toWrite string
		value := spill_manage.spill_thief()
		if value == 0 {
			toWrite = "ret\n"
		} else {
			toWrite = fmt.Sprintf("addq $%d, %%rsp\nret\n", value)
		}
		ValueWrite := fmt.Sprintf("###moving rsp back by %s###\n", value)
		c.writer.WriteString(ValueWrite)
		c.writer.WriteString(toWrite)
		c.writer.WriteString("###IN RETURN NODE###\n")
		//c.writer.WriteString("ret\n")
	case *LabelNode:
		c.writer.WriteString("###IN LABEL NODE###\n")
		newString := labelToASM(n.Label)
		//toWrite := fmt.Sprintf("%s\n", newString)
		//c.writer.WriteString(toWrite)
		return 0, "0", newString

	case *GotoNode:
		newString := labelToASM(n.Label)
		newString = fmt.Sprintf("jmp %s\n", strings.TrimRight(newString, ":"))
		c.writer.WriteString(newString)
		c.writer.WriteString("###It's a goto###\n")
		return 0, "nil", "nil"

	case *SysCallNode:
		tempHold := n.Label
		tempHoldFinal := strings.TrimRight(tempHold, ":")
		newStringFinal := fmt.Sprintf("call %s\n", tempHoldFinal)
		c.writer.WriteString(newStringFinal)
		//c.writer.WriteString("outyeah")
		return 0, "nil", "nil"

	case *CallNode:
		_, _, returnedString := c.compNode(n.Front()) //Change needed maybe
		toWrite := fmt.Sprintf("jmp %s\n%s\n", strings.TrimRight(returnedString, ":"), return_address)
		c.writer.WriteString(toWrite)
		return 0, "nil", "nil"

	case *TailcallNode:
		var toWrite string = ""
		var toWrite2 string

		_, _, lookup := c.compNode(n.Front()) //Change needed
		lookupVal := lookup[1 : len(lookup)-1]

		value := spill_manage.spill_thief()
		c.writer.WriteString("###CALLED SPILL_THIEF###\n")
		if value != 0 {
			toWrite = fmt.Sprintf("addq $%d, %%rsp\n", value)
		}

		if _, ok := allRegisters[lookupVal]; ok {
			toWrite2 = fmt.Sprintln("jmp *%s\n", lookup)
		} else {
			toWrite2 = fmt.Sprintf("jmp %s\n", strings.TrimRight(lookup, ":"))
		}

		stringToWrite := toWrite + toWrite2

		///

		///

		c.writer.WriteString(stringToWrite)

	case *TokenNode:
		//toWrite := fmt.Sprintf("TokenVal=%s\n", n.Value)
		//c.writer.WriteString(toWrite)
		//toWrite = fmt.Sprintf("TokenVal after change =%s\n", labelToASM(n.Value))
		//c.writer.WriteString("###IN TOKEN NODE###\n")
		return 0, "0", labelToASM(n.Value)

	case *AssignNode:
		var toInput2 string
		var toInputTemp string = ""
		switch child := n.Front().(type) {
		case *MemNode:

			_, _, reg := c.compNode(child.Front())
			offset := child.N8
			_ = n.Front()
			_, _, toInput := c.compNode(n.Next())
			if lastChar := toInput[0:1]; lastChar == "_" {

				toInputTemp = strings.TrimRight(toInput, ":")
				toInput2 = "$" + toInputTemp
			} else {
				toInput2 = toInput
			}

			toWrite := fmt.Sprintf("movq %s, %d(%s)\n", toInput2, offset, reg)
			c.writer.WriteString(toWrite)

			if toInput[0:1] == "_" {
				return_address = toInput
				c.writer.WriteString("subq $8, %rsp\n")
			}

			return 0, "nil", "nil"

		default:

			_, _, destination := c.compNode(n.Front())
			switch lChild := n.Next().(type) {

			default:
				_, _, inputToStore := c.compNode(lChild)
				if lastChar := inputToStore[:len(inputToStore)-1]; lastChar == ":" {

					inputToStore = strings.TrimRight(inputToStore, ":")
					inputToStore = "$" + inputToStore
				}
				toWrite := fmt.Sprintf("movq %s, %s\n", inputToStore, destination)
				c.writer.WriteString(toWrite)

				return 0, "nil", "nil"

			case *CmpopNode:
				bothNumb, val, _ := c.compNode(lChild) //*****
				_, _, rChild := c.compNode(n.Front())
				if bothNumb == 0 {
					toWrite := fmt.Sprintf("movq $%s, %s\n", val, rChild)
					c.writer.WriteString(toWrite)
				} else {
					eightMapped := "%" + eightBitEquiv[strings.TrimLeft(rChild, "%")]
					toWrite := fmt.Sprintf("%s %s\n", val, eightMapped)
					c.writer.WriteString(toWrite)
					toWrite = fmt.Sprintf("movzbq %s, %s\n", eightMapped, rChild)
					c.writer.WriteString(toWrite)
				}

				return 0, "nil", "nil"

			case *MemNode:

				memoryNode := lChild.Front()
				_, _, memX := c.compNode(memoryNode)
				memOffset := lChild.N8

				toWrite := fmt.Sprintf("movq %d(%s), %s\n", memOffset, memX, destination)
				c.writer.WriteString(toWrite)

				return 0, "nil", "nil"
			}

		}

	case *CmpopNode:
		op := n.Operator
		_, _, leftSide := c.compNode(n.Front()) //Change needed maybe
		_, _, rightSide := c.compNode(n.Next()) //Change needed maybe
		//toWrite2 := fmt.Sprintf("TokenVal rightSide after return =%s\n", leftSide)
		//c.writer.WriteString(toWrite2)
		leftSide = leftSide[1:]
		rightSide = rightSide[1:]

		//toWrite := fmt.Sprintf("if recursed 0: %d\n", numero)
		//c.writer.WriteString(toWrite)

		var opCode string

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

				var intToReturn string
				var boolToReturn bool
				switch op {

				default:
					return 0, "nil", "nil"

				case "<":
					boolToReturn = (ls < rs)
					//boolT := strconv.FormatBool(boolToReturn)
					//return 0, boolT, op

				case "<=":
					boolToReturn = (ls <= rs)
					//boolT := strconv.FormatBool(boolToReturn)
					//return 0, boolT, op

				case "=":
					boolToReturn = (ls == rs)
					//boolT := strconv.FormatBool(boolToReturn)
					//return 0, boolT, op
				}

				if boolToReturn {
					intToReturn = "1"
				} else {
					intToReturn = "0"
				}

				return 0, intToReturn, op

			} else {
				toWrite := fmt.Sprintf("cmpq $%d, %s\n", ls, labelToASM(rightSide))
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

				c.writer.WriteString("###In 2nd2 last Case###\n")
				return 1, opCode, op
			}
		} else {
			//if _, ok = strconv.Atoi(leftSide); ok == nil {
			//	leftSide = "$" + leftSide
			//}

			//toWrite := fmt.Sprintf("cmpq %s, %s\n", labelToASM(leftSide), labelToASM(rightSide))
			toWrite := fmt.Sprintf("cmpq %s, %s\n", labelToASM(rightSide), labelToASM(leftSide))
			c.writer.WriteString(toWrite)
			c.writer.WriteString("###In last Case###\n")
			toWrite = fmt.Sprintf("###leftSide is %s###\n###RightSide is %s###\n", leftSide, rightSide)
			c.writer.WriteString(toWrite)
			return 1, opCode, op
		}

	case *CjumpNode:
		_, _, op := c.compNode(n.Front())
		//yesLabel := n.Next().returnLabel()
		yesLabel := labelToASM(n.TrueLabel)
		//noLabel := n.Next().returnLabel()
		noLabel := labelToASM(n.FalseLabel)

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

		toWrite := fmt.Sprintf("%s %s\n", jmpStment, strings.TrimRight(yesLabel, ":"))
		c.writer.WriteString(toWrite)
		toWrite = fmt.Sprintf("jmp %s\n", strings.TrimRight(noLabel, ":"))
		c.writer.WriteString(toWrite)

	}
	return 0, "nil", "nil"
}

func (c *AsmCodeGenerator) BeginCompiler(ast Node) error {
	file, err := os.Create("prog.S")
	defer file.Close()
	if err != nil {
		return err
	}

	spill_manage.spillArray = make([]int, 1)
	spill_manage.Amount = 0
	spill_manage.spillArray[spill_manage.Amount] = 1
	c.writer = bufio.NewWriter(file)
	c.compNode(ast)
	c.writer.Flush()
	return nil
}
