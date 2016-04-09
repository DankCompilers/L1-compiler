package l1compiler

var count int

type Node interface {
	NodeId() int
	SetNodeId(int)
	AppendChild(Node)
	Next() Node
	Front() Node
	returnLabel() string
}

type ParseTreeNode struct {
	nodeId    int
	currChild int
	children  []Node
	Label     string
}

/// Helper function to make new Nodes
func newParseTreeNode(id int) ParseTreeNode {
	return ParseTreeNode{
		nodeId:   id,
		children: make([]Node, 0),
	}
}

func (p *ParseTreeNode) SetNodeId(id int) {
	p.nodeId = id
}

// Implementing Node interface

func (p *ParseTreeNode) NodeId() int {
	return p.nodeId
}

///
func (p *ParseTreeNode) Front() Node {
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
func (p *ParseTreeNode) returnLabel() string {
	return p.Label
}

func (p *OpNode) returnOp() string {
	return p.Operator
}

func (p *CmpopNode) returnCmpop() string {
	return p.Operator
}

///
func (p *ParseTreeNode) Next() Node {
	p.currChild += 1

	if p.currChild >= len(p.children) {
		return nil
	}

	return p.children[p.currChild]
}

/*********************** TYPE DEFINITIONS /********************************/

type ProgramNode struct {
	ParseTreeNode
	Label string
}

type SubProgramNode struct {
	ParseTreeNode
}

type FunctionNode struct {
	ParseTreeNode
	Label     string
	Arity     uint
	NumLocals uint
}

type InstructionNode struct {
	ParseTreeNode
}

type AssignNode struct {
	ParseTreeNode
}

type OpNode struct {
	ParseTreeNode
	Operator string
}

type CmpopNode struct {
	ParseTreeNode
	Operator string
}

type MemNode struct {
	ParseTreeNode
	X  string
	N8 uint
}

type LabelNode struct {
	ParseTreeNode
	Label string
}

type GotoNode struct {
	ParseTreeNode
	Label string
}

type CjumpNode struct {
	ParseTreeNode
	TrueLabel  string
	FalseLabel string
}

type SysCallNode struct {
	ParseTreeNode
	Label string
	Arity uint
}

type CallNode struct {
	ParseTreeNode
	Arity uint
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

func newProgramNode(func_label string, rest Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(rest)

	p := &ProgramNode{
		ParseTreeNode: b,
		Label:         func_label,
	}

	count++
	return p
}

func newSubProgramNode(funcNode, rest Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(funcNode)
	subcount := 0
	if rest != nil {
		b.AppendChild(rest)
		subcount++
	}

	p := &SubProgramNode{
		ParseTreeNode: b,
		//p.subcount = subcount
	}

	count++
	return p
}

func newFunctionNode(label string, arity int, locals int, rest Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(rest)

	p := &FunctionNode{
		ParseTreeNode: b,
		Label:         label,
		Arity:         uint(arity),
		NumLocals:     uint(locals),
	}

	count++
	return p
}

func newInstructionNode(instNode, rest Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(instNode)

	if rest != nil {
		b.AppendChild(rest)
	}

	p := &InstructionNode{
		ParseTreeNode: b,
	}

	count++
	return p
}

func newAssignNode(l, r Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(l)
	b.AppendChild(r)

	p := &AssignNode{
		ParseTreeNode: b,
	}

	count++
	return p
}

func newOpNode(op string, l, r Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(l)
	b.AppendChild(r)

	p := &OpNode{
		ParseTreeNode: b,
		Operator:      op,
	}

	count++
	return p
}

func newCmpopNode(op string, l, r Node) Node {
	b := newParseTreeNode(count)
	b.AppendChild(l)
	b.AppendChild(r)

	p := &CmpopNode{
		ParseTreeNode: b,
		Operator:      op,
	}

	count++
	return p
}

func newCjumpNode(cmpop Node, trueLabel, falseLabel string) Node {
	b := newParseTreeNode(count)
	b.AppendChild(cmpop)

	p := &CjumpNode{
		ParseTreeNode: b,
		TrueLabel:     trueLabel,
		FalseLabel:    falseLabel,
	}

	count++
	return p
}

func newMemNode(x Node, offset int) Node {
	b := newParseTreeNode(count)
	b.AppendChild(x)

	p := &MemNode{
		ParseTreeNode: b,
		N8:            uint(offset),
	}

	count++
	return p
}

func newLabelNode(label string) Node {
	b := newParseTreeNode(count)

	p := &LabelNode{
		ParseTreeNode: b,
		Label:         label,
	}

	count++
	return p
}

func newGotoNode(label string) Node {
	b := newParseTreeNode(count)

	p := &GotoNode{
		ParseTreeNode: b,
		Label:         label,
	}

	count++
	return p
}

func newSysCallNode(label string, arity int) Node {
	b := newParseTreeNode(count)

	p := &SysCallNode{
		ParseTreeNode: b,
		Label:         label,
		Arity:         uint(arity),
	}

	count++
	return p
}

func newCallNode(dest Node, arity int) Node {
	b := newParseTreeNode(count)
	b.AppendChild(dest)

	p := &CallNode{
		Arity: uint(arity),
	}

	count++
	return p
}

func newTailcallNode(dest Node, arity int) Node {
	p := &SysCallNode{
		Arity: uint(arity),
	}
	p.SetNodeId(count)
	p.AppendChild(dest)
	count++

	return p
}

func newReturnNode() Node {
	b := newParseTreeNode(count)

	p := &SysCallNode{
		ParseTreeNode: b,
	}

	count++
	return p
}

func newTokenNode(token string) Node {
	b := newParseTreeNode(count)

	p := &LabelNode{
		ParseTreeNode: b,
		Label:         token,
	}

	count++
	return p
}
