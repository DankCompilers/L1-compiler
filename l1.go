/(/		{ return LPAREN }
/)/		{ return RPAREN }
/;[a-zA-Z0-9_]*$/ { nlines++  /* ignore comments */ }
/[ \t]+/	{ /*ignore whitespace */ }

/:go/		{ return GOLABEL }
/return/	{ return RETURN }
/print/	{ return PRINT }
/allocate/	{ return ALLOCATE }
/array-error/	{ return ARRAYERROR }
/cjump/	{ return CJUMP }
/goto/	{ return GOTO }
/<-/	{ return ASSIGN }
/call/ { return CALL }
/tail-call/	{ return TAILCALL }

/\n/	{ nlines++ }

/[1-9][0-9]*/	{ return NAT }
/rcx/	{ return SCX }
/(rdi)|(rsi)|(rdx)|(r8)|(r9)		{  return A }
/(rax)|(rbx)|(rbp)|(r10)|(r11)|(r12|(r13)|(r14)|(r15)		{ return W}
/rsp/		{ return X }
/(<=?)|=/		{ return CMP }
/(<<=)|(>>=)	{ return SOP }
/(+=)|(-=)|(*=)|(&=)/	{ return AOP }
/^:[A-ZA-Z_][A-ZA-Z_0-9]*$/		{ return LABEL }
/./	{ fmt.Printf("Found invalid input on line: %d\n", nlines) }

//


package compiler

import ("fmt";"os")

func yylex(filename string) {
	nlines = 0
	lex := NewLexer(os.Stdin)
	txt = func() string { return lex.Text() }
	NNFun(lex)
}
