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

program: LPAREN LABEL subProgram RPAREN
{
	fmt.Printf("Detected program: %q\n", $2)
	$$ = newProgramNode($2, $3)
	cast(yylex).SetAstRoot($$)
}

subProgram: func
{
	fmt.Printf("Detected end subprogram\n")
	$$ = newSubProgramNode($1, nil)
}
| func subProgram
{
	fmt.Printf("Detected subprogram\n")
	$$ = newSubProgramNode($1, $2)
}


func: LPAREN LABEL NAT NAT subFunc RPAREN
{
	fmt.Printf("Detected func: %q\n", $2)
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


innerinstruction: w ASSIGN s
{
	fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $3)
}
|	w ASSIGN mem
{
	fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $3)
}
| mem ASSIGN s
{
	fmt.Printf("Detected assign \n")
	$$ = newAssignNode($1, $2)
}
| w aop t
{
	fmt.Printf("Detected aop \n")
	$$ = newOpNode($2, $1, $3)
}
| w sop sx
{
	fmt.Printf("Detected sop \n")
	$$ = newOpNode($2, $1, $3)
}
| w sop num
{
	fmt.Printf("Detected sop \n")
	$$ = newOpNode($2, $1, $3)
}
|  w ASSIGN cmp_op
{
	$$ = newAssignNode($1, $3)
}
|  GOTOLABEL LABEL
{
	fmt.Printf("Detected goto: %q\n", $2)
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
