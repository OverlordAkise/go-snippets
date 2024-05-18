# net/http logging without an extra module

To log a web request on your net/http webserver you can simply create a wrapper over your HandleFunc.

You prefix your HandleFunc function with a function that returns a HandlerFunc. Inside this returned function you run the actual web logic by manually calling it with "HandlerFunc(w,r)" , which means you can prefix this call with logging logic before and after.

I combined the logging logic into the defer function, because else it would have to be split between post-func and post-panic. It's better to log after the HandleFunc was called because then you can log the response code, request time and other stuff.
