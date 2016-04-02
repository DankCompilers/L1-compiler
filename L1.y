%{
package main

import (
		"io"
		"unicode/utf8"
		"fmt")

%}

%token CALLEEREG
%token CALLERREG
%token LABEL
%token STACKPNTR
%token NUM
%token AOP
%token SOP
%token CMP
%token CALL
%token LPAREN
%token RPAREN
%token GOLABEL
%token CJUMP
%token ALLOCATE
%token ASSIGN
%token TAILCALL
%token MEM
%token ARRAYERROR
%token PRINT
%token RETURN
%token RSP
%token RCX
%union { NUM int}

%%

expr: LPAREN subExp RPAREN
	|expr expr
;
subExp:
	  pPhrase {}
	| fPhrase {}
	| iPhrase {}
	| LPAREN subExp RPAREN {}

pPhrase:
	   LABEL fPhrase expr  {}
;

fPhrase:
	   LABEL NUM NUM expr {}
;

iPhrase:   CALLERREG ASSIGN s {}
		   CALLEEREG ASSIGN LPAREN MEM x NUM RPAREN {}
		|  LPAREN MEM x NUM RPAREN ASSIGN s  {}
		|  CALLEEREG AOP t  {}
		|  CALLERREG SOP sx  {}
		|  CALL ALLOCATE NUM {}
		|  CALL ARRAYERROR NUM {}
		|  CALL PRINT NUM {}
		|  TAILCALL u NUM {}
		|  RETURN {}
;

s	:  x {}
	| NUM {}
	| LABEL {}
;

t	: x {}
	| NUM {}
;

x	:CALLEEREG {}
    |RSP

;

u	:CALLEEREG {}
	|LABEL {}
;

sx: RCX {}
;
%%
