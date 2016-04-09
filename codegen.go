package l1compiler
//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"reflect"
// 	"regexp"
// 	"bytes"
// )
//
// var calleeArray = [...]string{"%rbx", "%rbp", "%r12", "%r13", "%r14", "%r15"}
// var callerArray = [...]string{"%rax", "%rcx", "%rdx", "%r8", "%r9", "%r10", "%r11"}
//
// var eightBitEquiv = map[string]string{
// 	"r10": "r10b",
// 	"r13": "r13b",
// 	"r8":  "r8b",
// 	"rbp": "bpl",
// 	"rdi": "dil",
// 	"r11": "r11b",
// 	"r14": "r14b",
// 	"r9":  "r9b",
// 	"rbx": "bl",
// 	"rdx": "dl",
// 	"r12": "r12b",
// 	"r15": "r15b",
// 	"rax": "al",
// 	"rcx": "cl",
// 	"rsi": "sil",
// }
//
// func labelToASM(label string) string {
// 	newLabel := strings.TrimLeft(label, ":")
// 	newLabel = "_" + newLabel
//
// 	return newLabel
// }
//
// type CodeGenerator interface {
// 	beginCompiler(Node) error
// }
//
// type AsmCodeGenerator struct {
//    seenLabels[string]string
//    seenFuncs map[string]FunctionNode
//    writer *bufio.Writer
// }
//
// func newASMGenerator() *AsmCodeGenerator {
//    return &AsmCodeGenerator{}
// }
//
//
// func genSaveCalleeRegisters() string {
//    var buff bytes.Buffer
//
//    for _, reg := range calleeArray {
// 	  toWrite := fmt.Sprintf("push  %%s \n", reg)
// 	  c.writer.WriteString(toWrite)
//    }
// }
//
// func (c *AsmCodeGenerator) compNode(node Node) {
// 	if node == nil {
// 		return
// 	}
//
// 	switch n := node.(type) {
// 	default:
// 		return
//
// 	case *ProgramNode:
// 	    //write header of code
// 		c.writer.WriteString(".text \n")
// 		c.writer.WriteString(".globl go \n")
// 		c.writer.WriteString("go: \n")
// 	   c.writer.WriteString(genSaveCalleeRegisters())
//
// 		c.compNode(n.children[0])
//
// 		for i := len(calleeArray) - 1; i >= 0; i-- {
// 			regToWrite := calleeArray[i]
// 			toWrite := fmt.Sprintf("pop  %%s \n", regToWrite)
// 			c.writer.WriteString(toWrite)
// 		}
// 		c.writer.WriteString("retq \n")
//
// 		//	case *AssignNode:
// 		//		mem, value := n.children[:]
// 		//		c.writer.WriteString("movq $%d, %s \n", value, mem)
//
// 	case *SubProgramNode, *InstructionNode:
// 		for index, element := range n.children {
// 			c.compNode(element)
// 		}
//
// 	case *FunctionNode:
// 		label := n.Label
// 		locals := n.NumLocals
// 		label = labelToASM(label)
// 		label = label + "\n"
// 		c.writer.WriteString(label)
// 		if locals > 0 {
// 			toWrite := fmt.Sprintf("subq $%d, %rsp", (8 * locals))
// 			c.writer.Writer(toWrite)
// 		}
// 		for index, element := range n.children {
// 			c.compNode(element)
// 		}
//
// 	case *OpNode:
// 		mem, value := n.children[:]
//
// 		if op := n.Operator; op == ">>=" {
// 			instruct := "sarq"
// 		} else {
// 			instruct := "salq"
// 		}
//
// 		if numRep, err := strconv.Atoi(value); err == nil {
// 			value = '$' + numRep
// 		}
//
// 		lowerReg = eightBitEquiv[mem]
// 		toWrite := fmt.Sprintf("%s %s, %s", instruct, value, lowerReg)
// 		c.writer.WriteString(toWrite)
//
// 	case *CallNode:
//
// 		u, NAT = n.children[:]
// 		u := u.value
// 		NAT := NAT
//
// 	case *LabelNode:
// 		newString := labelToASM(n.label)
// 		return newString
// 		//c.writer.WriteString(newString + "\n")
//
// 	case *GotoNode, *SysCallNode:
// 		newString := labelToASM(n.label)
// 		c.writer.WriteString("jmp %s", newString)
//
// 	case *CallNode:
// 		if arity := n.Arity; arity > 6 {
// 			spills := 8 * (arity - 6)
// 			toWrite = fmt.Sprintf("subq $%d %s", spills, "%rsp")
// 			c.writer.WriteString(toWrite)
// 		}
// 		toWrite = fmt.SprintF("jmp %s", c.compNode(n.children[0]))
// 		c.writer.WriteString(toWrite)
//
// 	case *AssignNode:
// 		if child := n.children[0]; child.(type) == *MemNode {
// 			reg := child.X
// 			offset := child.N8
// 			leftChildValue := n.children[1].token
//
// 			if i, err := strconv.Atoi(leftChildValue); !err {
// 				toInput := "$" + i
// 			} else {
// 				toInput := "%" + leftChildValue
// 			}
// 			toWrite := fmt.Sprintf("movq %s, %d(%s) \n", toInput, offset, reg)
// 			c.writer.Writer(toWrite)
// 		} else {
//
// 			destination = n.children[0].token
// 			if lChild := n.children[1]; lChild.(type) != *MemNode {
// 				tokenVal := lChild.token
// 				if _, ok := eightBitEquiv[tokenVal]; ok {
// 					//  c.writer.Writer("movq %%s, %%s", tokenVal, destination)
// 					inputToStore = "%" + tokenVal
//
// 				} else if i, err := strconv.Atoi(tokenVal); !err {
//
// 					inputToStore = "$" + tokenVal
// 				} else {
// 					inputToStore := c.compNode(lChild)
// 				}
//
// 				toWrite := fmt.Sprintf("movq %s, %s \n", inputToStore, destination)
// 				c.writer.WriteString(toWrite)
// 			} else if lChild.(type) == *CmpopNode {
// 				onleft, val := c.compNode(lChild)
// 				if onleft == 0 {
// 					rChild := n.children[0].token
// 					toWrite := fmt.Sprintf("movq $%s, %%s", rChild)
// 					c.writer.WriteString(toWrite)
// 				} else {
// 					eightMapped := eightBitEquiv[rChild]
// 					toWrite := fmt.Sprintf("%s %s", val, eightMapped)
// 					c.writer.WriteString(toWrite)
// 					toWrite = fmt.Sprintf("movzbq %s, %s", val, rChild)
// 					c.writer.WriteString(toWrite)
// 				}
//
// 			} else {
// 				memX := n.children[1].X
// 				memOffset := n.children.N8
//
// 				toWrite := fmt.Sprintf("movq %d(%%s), %s \n", memOffset, memX, destination)
// 				c.writer.WriteString(toWrite)
// 			}
//
// 		}
//
// 	case CmpopNode:
// 		op := Operator
// 		leftSide := n.children[0].token
// 		rightSide = n.children[1].token
//
// 		switch op {
// 		default:
// 			opCode := nil
// 		case "<":
// 			opCode := "setl"
// 		case "<=":
// 			opCode := "setge"
// 		case ">=":
// 			opCode := "setle"
// 		case "=":
// 			opCode := "sete"
// 		}
//
// 		if ls, ok := strconv.Atoi(leftSide); ok {
// 			if rs, ok := strconv.Atoi(rightSide); ok {
// 				switch op {
//
// 				default:
// 					return
//
// 				case "<":
// 					boolToReturn := (ls < rs)
// 					return 0, boolToReturn, op
//
// 				case "<=":
// 					boolToReturn := (ls <= rs)
// 					return 0, boolToReturn, op
//
// 				case "=":
// 					boolToReturn := (ls == rs)
// 					return 0, boolToReturn, op
// 				}
// 			} else {
// 				toWrite = fmt.Sprintf("cmpq $%d, %%s \n", ls, rightSide)
// 				c.writer.WriteString(toWrite)
// 				switch opCode {
// 				default:
// 					nil
// 				case "setl":
// 					opCode = "setg"
// 				case "setg":
// 					opCode = "setl"
// 				case "setge":
// 					opCode = "setle"
// 				case "setle":
// 					opCode = "setge"
// 				}
// 				return 1, opCode, op
// 			}
// 		} else {
// 			if _, ok = strconv.Atoi(leftSide); ok {
// 				leftSide = "$" + leftSide
// 			}
//
// 			toWrite = fmt.Sprintf("cmpq %s, %%s \n", leftSide, rightSide)
// 			c.writer.WriteString(toWrite)
// 			return 1, opCode, op
// 		}
//
// 	case *CjumpNode:
// 		yesLabel := n.children[1].label
// 		noLabel := n.children[2].label
// 		onleft, val, op = c.compNode(n.children[0])
// 		switch op {
// 		case "<=":
// 			jmpStment = "jle"
// 		case "<":
// 			jmpStment = "jl"
// 		case "=":
// 			jmpStment = "je"
// 		case ">=":
// 			jmpStment = "jge"
// 		}
// 		toWrite = fmt.Sprintf("%s %s", jmpStment, yesLabel)
// 		c.writer.WriteString(toWrite)
// 		toWrite = fmt.Sprintf("jmp %s", noLabel)
// 		c.writer.WriteString(toWrite)
//
// 	}
// }
//
// func (c *AsmCodeGenerator) genFromTree(ast Node, output_file string) error {
// 	file, err := os.Create(output_file)
// 	defer file.Close()
//
// 	if err != nil {
// 		fmt.Println("Could not open output file: ", err)
// 		return err
// 	}
//
// 	c.writer = bufio.NewWriter(file)
// 	c.compNode(ast)
// 	c.writer.Flush()
// 	return nil
// }
