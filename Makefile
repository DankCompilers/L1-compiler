UNAME := $(shell uname -s)

all:
	  ${GOPATH}/bin/nex lexer.nex

   # Could use nex instead of ed, but that'd be a little gratuitous.
   ifeq ($(UNAME), Darwin)
		 printf '/NEX_END_OF_LEXER_STRUCT/i\np *Compiler\n.\nw\nq\n' | ed -s lexer.nn.go
   else ifeq ($(UNAME), Linux)
			printf '/NEX_END_OF_LEXER_STRUCT/i\np *Compiler\n.\nw\nq\n' | ed -s lexer.nn.go
   else
			sed -i '/NEX_END_OF_LEXER_STRUCT/ip *Compiler' lexer.nn.go
   endif

	  go tool yacc -o=parser.yacc.go parser.y

test:
	go test
clean:
	-rm *.output *.yacc.go *.nn.go *.pdf *.dot L1generatedASM.txt
