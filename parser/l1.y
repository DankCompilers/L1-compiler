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
%type <node> program w u t INPUT sx NAT NUM LABEL MEM ASSIGN CMP RETURN TAILCALL subProgram
%type <node> func
%%

program: LPAREN ':' LABEL subProgram RPAREN {$$ =newProgrammingNode($3,$4)
										 cast(yylex).SetAstRoot}

subProgram: func {}
	| subProgram func {}

func:
	   LPAREN LABEL NAT NAT subFunc RPAREN {$$=newFunctionNode($2,$3,$4,$5)}
;

subFunc: inst {}
	| subFunc inst {}
;

inst: LPAREN innerinst RPAREN {$$=makeInstanceNode("Instance", $1,$2,$3)}

;

innerinst: ASSIGN_TO_MEM {$$=makeInnerinstNode($1)}

		   MEMTOREG {$$=makeInnerinstNode($1)}

		|  VALTOMEM {$$=makeInnerinstNode($1)}

		|  MEM_ASOP_INPUT {$$=makeInnerinstNode($1)}

		|  ASSIGN_CMP {$$=makeInnerinstNode($1)}

		|  LABEL {$$=makeInnerinstNode($1)}

		|  GOTOLABEL {$$=makeInnerinstNode($1)}

		|  SYSCALLS {$$=makeInnerinstNode($1)}

		|  JUMPCMP {$$=makeInnerinstNode($1)}

		|  TAIL {$$=makeInnerinstNode($1)}

		|  RETURN {$$=makeInnerinstNode($1)}

;

ASSIGN_CMP: w ASSIGN t CMP t {$$=makeASSIGN_CMPnode($1,$2,$3,$4,$5)}

;

TAIL: TAILCALL u NUM {$$=makeTailCallnode($1,$2,$3)}

;

JUMPCMP	 : CJUMP t CMP t LABEL LABEL {$$=makeJUMPCMPnode($1,$2,$3,$4,$5,$6)}

;

GOTOLABEL: GOTO LABEL {$$=makeGOTOnode($1,$2)}

;

ASSIGN_TO_MEM: w ASSIGN s {$$=makeASSIGNMEMnode($1,$2,$3)}

;

MEM_ASOP_INPUT: w ASOP INTOASOP {$$=makeMEMASOPINPUTnode($1,$2,$3)}

;

MEMTOREG:  w ASSIGN LPAREN MEM x NAT RPAREN {$$=makeMEMTOREGnode($1,$2,$3,$4,$5,$6,$7)}

;

VALTOMEM:  LPAREN MEM x NUM RPAREN ASSIGN s {$$=makeVALTOMEMnode($1,$2,$3,$4,$5,$6,$7)}

;

SYSCALLS: CALL SYSFUNC INPUT {$$makeCALLnode=($1,$2,$3)}

;

ASOP	: AOP {$$=makeTokenNode($1)}
		  SOP {$$=makeTokenNOde($1)}
;

INTOASOP: t {}
		| sx {}
		| NUM {$$=makeTokenNode($1)}
;

SYSFUNC: ALLOCATE {$$makeSYSFUNCnode($1)}
		|ARRAYERROR {$$makeSYSFUNCnode($1)}

		|PRINT {$$makeSYSFUNCnode($1)}

		|u {$$makeSYSFUNCnode($1)}

;

INPUT	:  NAT {$$=makeTokenNode($1)}
		|  '2' {$$=makeTokenNode($1)}
		|  '1' {$$=makeTokenNode($1)}
;

u	: w {}
	| LABEL {$$=makeTokenNode($1)}
;

t	: x {}
	| num {}
;

s	:  x {}
	| num {}
	| LABEL {$$=makeTokenNode($1)}
;

x	: RSP {$$=makeTokenNode($1)}
	| w   {}
;

w	: W {}
	| a {}
;

a	: ARG {$$=makeTokenNode($1)}
	| sx {}
;

sx	: RCX {$$=makeTokenNode($1)}

;

num	: 	NUM {$$=makeTokenNode($1)}
	|	NAT {$$=makeTokenNode($1)}
%%
