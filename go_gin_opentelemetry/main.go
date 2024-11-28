package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//database
	db "luctus.at/example/db"
	//tracing
	"context"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	//tracing - unknown_service fix
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	//for getting traceid
	//tr "go.opentelemetry.io/otel/trace"
	//This only adds the log as an event in the span, not important:
	//"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func SetupTracer() (*trace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("otel-example"),
			semconv.ServiceVersion("1.0.0"),
			semconv.DeploymentEnvironment("dev"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func main() {
	//Setup tracing
	tp, err := SetupTracer()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	db.InitDatabase(tp)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(otelgin.Middleware("istinaweb-gin"))

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/example", func(c *gin.Context) {
		//instead of c.HTML you use otelgin.HTML
		//the first argument is the context, the rest is the same
		otelgin.HTML(c, 200, "example", gin.H{
			"title": "Example page",
		})
	})

	r.GET("/sql", func(c *gin.Context) {
		otelgin.HTML(c, 200, "sql", gin.H{
			"title": "Example sql page",
			"list":  db.GetColumns(c.Request.Context()),
		})
	})

	fmt.Println("Now listening on :8832")
	fmt.Println(r.Run("0.0.0.0:8832"))
}
