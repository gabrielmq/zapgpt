package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielmq/zapgpt/internal/handlers"
)

func main() {
	lambda.Start(handlers.Handle)
}
