# WSDLs in golang

To generate go code from .wsdl files I used the following library:

https://github.com/hooklift/gowsdl

## Install wsdl2go generator

    wget https://github.com/hooklift/gowsdl/archive/refs/heads/master.zip
    unzip master.zip
    cd gowsdl-master
    go get .
    go build cmd/gowsdl/main.go
    sudo mv main /usr/local/bin/gowsdl

Then simply:

    gowsdl -o edl.go -p edl wsdls/my.wsdl

## Info

The gowsdl command generates a <name>.go and a server<name>.go file.  
The serverXXXX.go file actually contains a method called "Endpoint". This can be used as a net/http listener function like this:

    http.HandleFunc("/wsdl/",edl.Endpoint)

Now you just have to fill out the few one-line functions in the serverXXXX.go file so that they contain logic.

You could also use your own framework and simply c.Bind() the incoming data to the above generated wsdl objects.
