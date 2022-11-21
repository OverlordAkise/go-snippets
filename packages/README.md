# Using packages to manage complexity

At the beginning you may simply have a file for every "module" of your application. Example: 

 - main.go
 - db.go
 - logger.go

You can use packages to put specific "modules" of your application into sorted, modular and reusable packages.

An example is in this folder:

 - main.go
 - database/mysql.go
 - logging/logger.go

This way you can have a lot of files in your application without your folders or code getting cluttered.

_**VERY IMPORTANT:**_  
Lets say you initialized your go mod file with the following:

    go mod init example.com/myapp

You have to change the first line with "package" in your extra package-folder to something like:

    package logger

And then in your main.go (or wherever you want to use it) you should import it like this:

    import example.com/myapp/logger

Without this it wont import properly!
