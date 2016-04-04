%{
package l1compiler

import (
		"fmt"
		"strconv"
)

%}

%union {
	n int
	s string
	node Node
}


%token <s> LABEL GOLABEL
%token <n> NEG NAT NAT6 N8 NEGN8
%token <s> AOP SOP CMP
%token CALL CJUMP TAILCALL RETURN GOTO
%token LPAREN RPAREN
%token <s> ALLOCATE PRINT ARRAYERROR
%token ASSIGN MEM
%token <s> RSP RCX
%token <s> X W A
%type <node> program subProgram func subFunc instruction innerinstruction mem cmp_op syscalls s w u t x a sx num label
%type <n> nat n8
%%

program: LPAREN LABEL subProgram RPAREN
{
	fmt.Println("Detected program: %+v", $2)
	$$ = newProgramNode($2, $3)
	cast(yylex).SetAstRoot($$)
}

subProgram: func
{
	fmt.Println("Detected end subprogram")
	$$ = newSubProgramNode($1, nil)
}
| func subProgram
{
	fmt.Printf("Detected subprogram\n")
	$$ = newSubProgramNode($1, $2)
}


func: LPAREN LABEL nat nat subFunc RPAREN
{
	fmt.Println("Detected func: %+v", $2)
	$$ = newFunctionNode($2, $3, $4, $5)
}


subFunc: instruction
{
	fmt.Printf("Detected end instruction\n")
	$$ = newInstructionNode($1, nil)
}
|  instruction subFunc
{
	fmt.Printf("Detected instruction\n")
	$$ = newInstructionNode($1, $2)
}


instruction: LPAREN innerinstruction RPAREN
{
	$$ = $2
}
|  label
{
	$$ = $1
}


innerinstruction: w ASSIGN s
{
	//fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $3)
}
|	w ASSIGN mem
{
	//fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $3)
}
| mem ASSIGN s
{
	//fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $3)
}
| w AOP t
{
	//fmt.Printf("Detected aop \n")
	$$ = newOpNode($2, $1, $3)
}
| w SOP sx
{
	//fmt.Printf("Detected sop \n")
	$$ = newOpNode($2, $1, $3)
}
| w SOP num
{
	//fmt.Printf("Detected sop \n")
	$$ = newOpNode($2, $1, $3)
}
|  w ASSIGN cmp_op
{
	$$ = newAssignNode($1, $3)
}
|  GOTO LABEL
{
	//fmt.Printf("Detected goto: %q\n", $2)
	$$ = newGotoNode($2)
}
| CJUMP cmp_op LABEL LABEL
{
	$$ = newCjumpNode($2, $3, $4)
}
|  syscalls
{
	$$ = $1
}
| CALL u nat
{
	$$ = newCallNode($2, $3)
}
|  TAILCALL u NAT6
{
	$$ = newTailcallNode($2, $3)
}
|  RETURN
{
	$$ = newReturnNode()
}




syscalls: CALL PRINT NAT6
{
	$$ = newSysCallNode($2, 1)
}
| CALL ALLOCATE NAT6
{
	$$ = newSysCallNode($2, 2)
}
| CALL ARRAYERROR NAT6
{
	$$= newSysCallNode($2, 3)
}


cmp_op: t CMP t
{
	$$ = newCmpopNode($2, $1, $3)
}


mem: LPAREN MEM x n8 RPAREN
{
	$$ = newMemNode($3, $4)
}




u: 	w 		{ $$ = $1 }
	| label { $$ = $1	}


t: 	x 	{ $$ = $1 }
	| num { $$ = $1 }


s:  x  { $$ = $1 }
	| num { $$ = $1 }
	| label { $$ = $1 }

x: X	 { $$ = newTokenNode($1) }
	| w   { $$ = $1 }


w: W { $$ = newTokenNode($1) }
	| a { $$ = $1 }


a: A { $$ = newTokenNode($1) }
	| sx { $$ = $1 }


sx: RCX { $$ = newTokenNode($1) }

num: 	NAT { $$ = newTokenNode(strconv.Itoa($1)) }
|	NEG 		{ $$ = newTokenNode(strconv.Itoa($1)) }
| n8			{ $$ = newTokenNode(strconv.Itoa($1)) }


nat: NAT 	{ $$ = $1 }
| NAT6 		{ $$ = $1 }
|	N8 			{ $$ = $1 }

n8: N8 	{ $$ = $1 }
| NEGN8 { $$ = $1 }
| NAT6	{ $$ = $1 }


label: LABEL
{
	$$ = newLabelNode($1)
}
%%
func cast(y yyLexer) *Compiler { return y.(*Lexer).p }
