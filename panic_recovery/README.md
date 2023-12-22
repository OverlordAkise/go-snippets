# Goroutine panic and recover

This is a simple test that shows:

If a goroutine panics it will crash the whole go execution.  
This means catching panics of goroutines is important!

Catching a panic can be done via a "recover" method in a "defer" function.  
Defer runs after the current function ended, even during a panic, and a call to recover stops a panic from "bubbling further up".

This is important for database transactions.  
If you begin a transaction in sql and a panic happens then the connection stays open and in use (and the transaction too).  
This means you **have to** have a defer and recover function that atleast calls `tx.Rollback` so that the transaction closes and the connection too.

