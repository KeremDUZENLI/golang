package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(ctx context.Context, r Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", r.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
