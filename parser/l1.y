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

innerinst: w ASSIGN s {}
		|  w ASSIGN LPAREN MEM x NAT RPAREN {}
		|  LPAREN MEM x NUM RPAREN ASSIGN s  {}
		|  w ASOP INTOASOP {}
		|  w ASSIGN t CMP t {}
		|  LABEL {}
		|  GOTO LABEL {}
		|  CJUMP t CMP t LABEL LABEL {}
		|  SYSCALLS {}
		|  TAILCALL u NUM {}
		|  RETURN {}
;

SYSCALLS: CALL SYSFUNC INPUT {}

;

ASOP: AOP
	  SOP
;

INTOASOP: t
		| sx
		| NUM
;

SYSFUNC: ALLOCATE {}
		|ARRAYERROR {}
		|PRINT {}
		|u {}
;

INPUT:  NAT {}
		'2' {}
		'1' {}
;

u: w {}
	| LABEL {}
;

t: x {}
	| num {}
;

s	:  x {}
	| num {}
	| LABEL {}
;

x	: X {}
	| w {}
;

w:  W {}
	| a {}


a: ARG {}
	| sx {}
;

sx: RCX {}
;

num: 	NUM {}
	|	NAT {}
%%
