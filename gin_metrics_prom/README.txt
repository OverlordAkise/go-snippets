# Prometheus metrics in gin-gonic

This example shows:

 - How to enable monitoring for gin-gonic
 - How to add new metrics
 - How to expose /metrics in gin routes

To test this:

 - Start the server with `go run .`
 - Visit http://localhost:8080/ and http://localhost:8080/panic
 - Visit http://localhost:8080/metrics

For your prometheus config (in this example) enter `http://localhost:8080` as your target.
