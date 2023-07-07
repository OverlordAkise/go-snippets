# Rotating logs with zap logging on a time basis

To do this we use the logrotate service on linux servers.

If you are on kubernetes or any cloud infrastructure it is better to log to stdout and handle the logs that way.  

On Linux it is easier to handle log rotations outside of your application, aka. "Separation of concerns".  
Because logging is different on different infrastructures I think its a nice pattern to not have a rotating logging solution built into your application.

## Go application

The main.go file creates a simple zap logger with its output path set to myapp.log in the current directory.

## Configure logrotate.d

We now want to rotate this myapp.log file on a daily basis.  
We assume the file lies in /var/log/myapp/access.log

You can create a file (as root) in /etc/logrotate.d/myapp with the following content:

```
/var/log/myapp/*.log {
        daily
        missingok
        rotate 14
        compress
        delaycompress
        notifempty
        create 0640 user group
}
```

This will now automatically rotate your logs and compress them every day at midnight.  
For other examples you can look at /etc/logrotate.d/nginx , as nginx is rotating its logs with logrotate too.
