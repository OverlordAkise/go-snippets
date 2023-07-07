# Using OpenTelemetry with gin and sql

This example shows:

 - How to setup tracing in golang
 - How to trace gin (specifically HTML rendering)
 - How to trace sql (with example database/sql)

This was tested with Grafana Tempo and the Grafana WebUI.

# Explanation

In this example a trace starts when a HTTP request comes in.

We then propagate the context, where the trace info is inside of, to the sql package where the "otelsql" adds it's span for the db.Query. The otelsql package automatically measures any sql function and also has an extra version for the "sqlx" package.

Then the gin HTML renderer starts and this will automatically, because we use otelgin.HTML instead of c.HTML, create another span in our trace.

The traces are cached and will be periodically sent out to the backend server. This example "samples" every trace, so that you can see it immediately in the grafana ui.
