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

%%

program: LPAREN LABEL subProgram RPAREN

subProgram: func
	| subProgram func

func:
	   LPAREN LABEL NAT NAT subFunc RPAREN {}
;

subFunc: inst {}
	| subFunc inst {}
;

inst: LPAREN innerinst RPAREN {}

innerinst:   w ASSIGN s {}
		|  w ASSIGN LPAREN MEM x NAT RPAREN {}
		|  LPAREN MEM x NUM RPAREN ASSIGN s  {}
		|  w AOP t  {}
		|  w SOP sx  {}
		|  w ASSIGN t CMP t {}
		|  LABEL {}
		|	 GOTO LABEL {}
		|  CJUMP t CMP t LABEL LABEL {}
		|  CALL u NAT
		|  CALL ALLOCATE '2' {}
		|  CALL ARRAYERROR '2' {}
		|  CALL PRINT '1' {}
		|  TAILCALL u NUM {}
		|  RETURN {}
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
	|		NAT {}
%%
