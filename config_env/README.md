# Config with environment variables

If you run this example like this it will say "empty":

```bash
go run main.go
```

You can run it like this to set it temporarily:

```bash
MYAPP_PORT=9000 go run .
```

Or like this for "longevity":

```bash
export MYAPP_PORT=9000
go run .
```
