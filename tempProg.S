.text
.globl go

go:
push  %rbx
push  %rbp
push  %r12
push  %r13
push  %r14
push  %r15
call _printlocal



pop %r15
pop %r14
pop %r13
pop %r12
pop %rbp
pop %rbx
retq
_printlocal:
subq $8, %rsp 
###CALLING SPILL ACCRUE###
###IN FUNCTION NODE###
movq $11, %rdi
movq %rdi, 0(%rsp)
movq $7, %rdi
###IN LABEL NODE###
movq $_f_ret, -8(%rsp)
subq $8, %rsp
###IN LABEL NODE###
jmp _f
_f_ret:
###IN LABEL NODE###
_f_ret:
movq 0(%rsp), %rdi
call print
###moving rsp back by %!s(int=8)###
addq $8, %rsp
ret
###IN RETURN NODE###
_f:
subq $24, %rsp 
###CALLING SPILL ACCRUE###
###IN FUNCTION NODE###
movq %rdi, 0(%rsp)
movq $3, 8(%rsp)
movq $5, 16(%rsp)
###moving rsp back by %!s(int=24)###
addq $24, %rsp
ret
###IN RETURN NODE###



