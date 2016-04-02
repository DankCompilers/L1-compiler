%{
package main

/*import (
		"io"
		"unicode/utf8"
		"fmt"
)*/

%}

%union {
	n int
	s string
}


%token LABEL GOLABEL
%token NUM NAT
%token AOP SOP CMP
%token CALL CJUMP TAILCALL RETURN GOTO
%token LPAREN RPAREN
%token ALLOCATE PRINT ARRAYERROR
%token ASSIGN MEM
%token RSP RCX
%token X W ARG
%type <node> program w u t INPUT sx NAT NUM LABEL MEM ASSIGN CMP RETURN TAILCALL
%%

program: LPAREN LABEL subProgram RPAREN {$$ =newProgrammingNode($2,$3)
										 cast(yylex).SetAstRoot}

subProgram: func
	| subProgram func

func:
	   LPAREN LABEL NAT NAT subFunc RPAREN {$$=newFunctionNode($2,)}
;

subFunc: inst {}
	| subFunc inst {}
;

inst: LPAREN innerinst RPAREN {}

innerinst: ASSIGN_TO_MEM {}
		   MEMTOREG {}
		|  VALTOMEM {}
		|  MEM_ASOP_INPUT {}
		|  w ASSIGN t CMP t {}
		|  LABEL {}
		|  GOTOLABEL {}
		|  SYSCALLS {}
		|  JUMPCMP {}
		|  TAIL {}
		|  RETURN {}
;

ASSIGN_CMP: w ASSIGN t CMP t {}

;

TAIL: TAILCALL u NUM {}

;

JUMPCMP	 : CJUMP t CMP t LABEL LABEL {}

;

GOTOLABEL: GOTO LABEL {}

;

ASSIGN_TO_MEM: w ASSIGN s {}

;

MEM_ASOP_INPUT: w ASOP INTOASOP {}

;

MEMTOREG:  w ASSIGN LPAREN MEM x NAT RPAREN {}

;

VALTOMEM:  LPAREN MEM x NUM RPAREN ASSIGN s {}

;

SYSCALLS: CALL SYSFUNC INPUT {}

;

ASOP: AOP {}
	  SOP {}
;

INTOASOP: t {}
		| sx {}
		| NUM {}
;

SYSFUNC: ALLOCATE {}
		|ARRAYERROR {}
		|PRINT {}
		|u {}
;

INPUT	:  NAT {}
		|  '2' {}
		|  '1' {}
;

u	: w {}
	| LABEL {}
;

t	: x {}
	| num {}
;

s	:  x {}
	| num {}
	| LABEL {}
;

x	: X {}
	| w {}
;

w	: W {}
	| a {}
;

a	: ARG {}
	| sx {}
;

sx	: RCX {}

;

num	: 	NUM {}
	|	NAT {}
%%
