package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/events"
)

func returnIP(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProyResponse {
    StatusCode: 200,
    Body: "127.0.0.7",
  }, nil
}

func main() {
  lambda.Start(returnIP)
}
