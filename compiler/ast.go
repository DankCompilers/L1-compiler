package L1

type Node interface {
	AppendChild(Node)
}

type ParseTreeNode struct {
	NodeID   string
	children []Node
}

func makeParseTreeNode(str string) Node {
}

func makeTokenNode(str string) Node {
	b := MakePareTreeNode()
}

func makeSYSFUNCnode(str string) Node {
}

func makeCALLnode(str string) Node {

}

func makeVALTOMEMnode(mem Node, str2,str3,str4,str5,str6) Node {
	return 0
}

func makeMEMTOREGnode() Node {

}

func makeMEMASOPINPUTnode() Node {

}

func makeASSINGNMENTnode() Node {

}

func makeGOTOnode() Node {

}

func makeJUMPCMPnode () Node {
}

func makeTailCallNode() Node {
}

func makeASSIGN_CMPnode() Node {

}

func makeInnerinstNode() Node {

}

func makeInstanceNode() Node {

}

func newFunctionNode() Node {

}
