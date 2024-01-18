package main

import (
	"context"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
)

const pspanKey = 42

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	span := opentracing.StartSpan("helloHandler")
	defer span.Finish()

	span.LogKV(
		"handler", "this is handler log",
	)
	span.SetTag("handler", "this is handler tag")

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	helloService(ctx)
}

func helloService(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "helloService")
	defer span.Finish()

	span.LogKV(
		"service", "this is service log",
	)

	span.SetTag("service", "this is service tag")

	helloDatabase(ctx)
}

func helloDatabase(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "helloDatabase")
	defer span.Finish()

	span.LogKV(
		"database", "this is database log",
	)

	span.SetTag("database", "this is database tag")

	time.Sleep(time.Second)
}
