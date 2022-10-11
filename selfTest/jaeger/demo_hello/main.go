package main

import (
	"fmt"
	"os"

	"geekGo/selfTest/jaeger/lib/tracing"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	if len(os.Args) != 2 {
		panic("ERROR: Expecting one argument")
	}

	tracer, closer := tracing.Init("hello-world","127.0.0.1:6831")
	defer closer.Close()

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)

	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)

	println(helloStr)
	span.LogKV("event", "println")

	span.Finish()
}



