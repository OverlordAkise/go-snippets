# Goroutine panic tests

This is a simple test that shows:

If a goroutine panics it will crash the whole go execution. 

This means catching panics of goroutines is important!

Catching a panic can be done via a "recover" method in a "defer" function.

But: A defer runs before the panic even if a panic happens in the current function. (example: no_recover.go)

Defer runs after the current function ended, and a call to recover stops a panic and (should) handle it.

