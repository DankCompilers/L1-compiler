%{
package L1

/*import (
		"io"
		"unicode/utf8"
		"fmt"
		"strconv"
)*/

%}

%union {
	s string
	node Node
}


%token <s> LABEL GOLABEL
%token <s> NEG NAT NAT6 NAT8
%token <s> AOP SOP CMP
%token CALL CJUMP TAILCALL RETURN GOTO
%token LPAREN RPAREN
%token <s> ALLOCATE PRINT ARRAYERROR
%token ASSIGN MEM
%token <s> RSP RCX
%token <s> X W A
%type <node> program w u t sx num NAT NAT6 NAT8 NEG label MEM ASSIGN CMP RETURN TAILCALL subProgram
%type <node> func
%%

program: LPAREN GOLABEL subProgram RPAREN
{
	$$ = newProgramNode($2,$3)
	cast(yylex).SetAstRoot($$)
}

subProgram: func
{
	$$ = newSubProgramNode($1, nil)
}
| func subProgram
{
	$$ = newSubProgramNode($1, $2)
}


func: LPAREN LABEL NAT NAT subFunc RPAREN
{
	$$ = newFunctionNode($2, $3, $4, $5)
}


subFunc: instruction
{
	$$ = newInstructionNode($1, nil)
}
|  instruction subFunc
{
	$$ = newInstructionNode($1, $2)
}


instruction: LPAREN innerinstruction RPAREN
{
	$$ = $2
}


innerinstruction: w ASSIGN s
{
	$$ = newAssignNode($1, $3)
}
|	w ASSIGN mem
{
	$$ = newAssignNode($1, $3)
}
| mem ASSIGN s
{
	$$ = newAssignNode($1, $2)
}
| w aop t
{
	$$ = newOpNode($2, $1, $3)
}
| w sop sx
{
	$$ = newOpNode($2, $1, $3)
}
| w sop num
{
	$$ = newOpNode($2, $1, $3)
}
|  w ASSIGN cmp_op
{
	$$ = newAssignNode($1, $3)
}
|  GOTOLABEL LABEL
{
	$$ = newGotoNode($2)
}
| CJUMP cmp_op label label
{
	$$ = newCjumpNode($2, $3, $4)
}
|  syscalls
{
	$$ = $1
}
| CALL u NAT
{
	$$ = newCallNode($2, $3)
}
|  tailcall
{
	$$ = $1
}
|  RETURN
{
	$$ = newReturnNode()
}
|  label
{
	$$ = $1
}



tailcall: TAILCALL u NAT6
{
	$$ = newTailcallNode($2, $3)
}


SYSCALLS: CALL PRINT '1'
{
	$$ = newSysCallNode($2, $3)
}
| CALL ALLOCATE '2'
{
	$$ = newSysCallNode($2, $3)
}
| CALL ARRAYERROR '2'
{
	$$= newSysCallNode($2, $3)
}


cmp_op: t CMP t
{
	$$ = newCmpopNode($2, $1, $3)
}


mem: LPAREN MEM x NAT8 RPAREN
{
	$$ = newMemNode($2, )
}


aop: AOP
{
	$$ = $1
}


sop: SOP
{
	$$ = $1
}

u: 	w 		{ $$ = $1 }
	| label { $$ = $1	}


t: 	x 	{ $$ = $1 }
	| num { $$ = $1 }


s:  x  { $$ = $1 }
	| num { $$ = $1 }
	| label { $$ = $1 }

x: RSP { $$ = newTokenNode($1) }
	| w   { $$ = $1 }


w: W { $$ = newTokenNode($1) }
	| a { $$ = $1 }


a: A { $$ = newTokenNode($1) }
	| sx { $$ = $1 }


sx: RCX { $$ = newTokenNode($1) }


num: 	NAT { $$ = newTokenNode(strconv.Itoa($1)) }
	|	NEG 	{ $$ = newTokenNode(strconv.Itoa($1)) }
	| NAT6 	{ $$ = newTokenNode(strconv.Itoa($1)) }
	| NAT8 	{ $$ = newTokenNode(strconv.Itoa($1)) }


label: LABEL
{
	$$ = newLabelNode($1)
}
%%
func cast(y yyLexer) *Compiler { return y.(*Lexer).p }
