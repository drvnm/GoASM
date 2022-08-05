format ELF64 executable
segment readable executable
entry start
start:
    mov rax, 1
    mov rdi, 1
    mov rsi, hello
    mov rdx, 14
    syscall

    mov rax, 60
    mov rdi, 0
    syscall

segment readable writable
hello: db "Hello, World!", 10