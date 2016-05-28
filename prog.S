.text
.globl go

go:
push  %rbx
push  %rbp
push  %r12
push  %r13
push  %r14
push  %r15
call _go



pop %r15
pop %r14
pop %r13
pop %r12
pop %rbp
pop %rbx
retq
_go:
###CALLING FUNCTION NODE WITH SPILL: 0###
###IN FUNCTION NODE###
movq $5, %rdi
movq $5, %rsi
call allocate
movq %rax, %rdi
call print
###moving rsp back by %!s(int=0)###
ret
###IN RETURN NODE###


