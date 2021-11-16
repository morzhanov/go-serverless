package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/morzhanov/go-serverless/internal"
)

var s = internal.NewService()

func HandleRequest(_ context.Context, ev internal.Event) {
	b, _ := json.Marshal(ev)
	s.Handle(b)
}

func main() {
	lambda.Start(HandleRequest)
}
