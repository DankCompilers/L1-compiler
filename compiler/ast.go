package L1

var count int


type Node interface {
	NodeId() int
	AppendChild(Node)
	Next() Node
	Front() Node
}

type ParseTreeNode struct {
	nodeId   	int
	currChild int
	children 	[]Node
}

/// Helper function to make new Nodes
func newParseTreeNode(id int) ParseTreeNode {
	return ParseTreeNode {
		nodeId: id,
		children : new([]Node, 0),
	}
}


// Implementing Node interface

func (p *ParseTreeNode) NodeId() int {
	return p.nodeId
}

///
func (p *ParseTreeNode) Front() ParseTreeNode {
	p.currChild = 0

	if len(p.children) <= 0 {
		return nil
	}

	return p.children[p.currChild]
}

///
func (p *ParseTreeNode) AppendChild(n Node) {
	p.children = append(p.children, n)
}

///
func (p *ParseTreeNode) Next() {
	p.currChild += 1

	if p.currChild >= len(p.children) {
		return nil
	}

	return p.children[p.currChild]
}


/*********************** TYPE DEFINITIONS /********************************/

type ProgramNode struct {
	ParseTreeNode
	EntryLabel string
}

type SubProgramNode struct {
	ParseTreeNode
}

type FunctionNode struct {
	ParseTreeNode
	Label 				string
	Arity 				uint
	NumLocals 		uint
}

type InstructionNode struct {
	ParseTreeNode
}

type AssignNode struct {
	ParseTreeNode
}

type OpNode struct {
	ParseTreeNode
	Operator 			string
}

type CmpopNode struct {
	ParseTreeNode
	Operator 			string
}

type MemNode struct {
	ParseTreeNode
	X 						string
	N8 						uint8
}

type LabelNode struct {
	ParseTreeNode
	Label 				string
}

type GotoNode struct {
	ParseTreeNode
	Label 				string
}

type SysCallNode struct {
	ParseTreeNode
	Label 				string
	Arity					uint
}

type CallNode struct {
	ParseTreeNode
	Arity 				uint
}


type TailcallNode struct {
	ParseTreeNode
}

type ReturnNode struct {
	ParseTreeNode
}

type TokenNode struct {
	ParseTreeNode
	Value string
}


/*********************** HELPER DEFINITIONS /********************************/


func newProgramNode(func_label string, rest Node) ProgramNode {
	p := &ProgramNode {
		nodeId	: count,
		Label		: func_label
	}
	p.AppendChild(rest)
	count++

	return p
}


func newSubProgramNode(funcNode, rest Node) SubProgramNode {
	p := &SubProgramNode{ nodeId: count }
	p.AppendChild(funcNode)
	p.AppendChild(rest)
	count++

	return p
}


func newFunctionNode(label string, arity int, locals int, rest Node) FunctionNode {
	p := &FunctionNode{
		nodeId: count,
		Label: label,
		Arity: uint(arity),
		NumLocals: uint(locals)
	}

	p.AppendChild(rest)
	count++

	return p
}


func newInstructionNode(instNode Node, rest InstructionNode) InstructionNode {
	p := &InstructionNode{ nodeId: count }
	p.AppendChild(instNode)
	p.AppendChild(rest)
	count++

	return p
}


func newAssignNode(l, r Node) AssignNode {
	p := &AssignNode{ nodeId: count }
	p.AppendChild(l)
	p.AppendChild(r)
	count++

	return p
}


func newOpNode(op string, l, r Node) OpNode {
	p := &OpNode{
		nodeId: count,
		Operator: op
	}

	p.AppendChild(l)
	p.AppendChild(r)
	count++

	return p
}


func newCmpopNode(op string, l, r Node) CmpopNode {
	p := &CmpopNode{
		nodeId: count,
		Operator: op
	}

	p.AppendChild(l)
	p.AppendChild(r)
	count++

	return p
}


func newMemNode(x TokenNode, offset int) MemNode {
	p := &MemNode{
		nodeId: count,
		X: x.Value,
		N8: uint(offset)
	}

	count++

	return p
}


func newLabelNode(label string) LabelNode {
	return &LabelNode{
		nodeId	: count++,
		Label		: label
	}
}


func newGotoNode(label string) GotoNode {
	p := &GotoNode{
		nodeId: count,
		Label: label
	}
	count++

	return p
}


func newSysCallNode(label string, arity int) SysCallNode {
	p := &SysCallNode{
		nodeId: count,
		Label: label,
		Arity: uint(arity),
	}
	count++

	return p
}


func newCallNode(dest Node, arity int) Node {
	p := &CallNode{
		nodeId: count,
		Arity: uint(arity),
	}

	p.AppendChild(dest)
	count++

	return p
}


func newTailcallNode(dest Node, arity int) TailcallNode {
	p := SysCallNode{
		nodeId: count,
		Arity: uint(arity)
	}
	p.AppendChild(dest)
	count++

	return p
}


func newReturnNode() Node {
	p := SysCallNode{ nodeId: count }
	count++

	return p
}

func newTokenNode(token string) Node {
	return LabelNode{
		nodeId	: count++,
		Value		: token
	}
}
