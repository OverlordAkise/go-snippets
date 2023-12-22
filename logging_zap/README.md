# Logging in golang (and gin) with zap 

Nowadays you log everything in JSON format so that a log server can parse and interpret these logs easily.

A very popular solution for logging is zap.  
It is very fast, easy to use and has a lot of "community support" because of its popularity.

The `main.go` file shows the basics of using zap.  
The `gin.go` file shows how to use zap and "ginzap" to log web requests.  
The `rotation.go` file shows how to rotate files with zap.

