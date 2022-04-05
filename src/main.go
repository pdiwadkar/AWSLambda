package main

/*
aws ID-  <7681-9119-0495>
aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
aws lambda create-function --function-name goaws --zip-file fileb://main.zip --handler main --runtime go1.x --role arn:aws:iam::768191190495:role/lambda-ex
*/
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyRequest struct {
	Name string `json:"what is your name?"`
	Age  int    `json:"what is your age?"`
}
type MyResponse struct {
	Message string `json:"Answer"`
}

func FirstLambdaFunction(evt MyRequest) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", evt.Name, evt.Age)}, nil
}

func main() {
	lambda.Start(FirstLambdaFunction)

}
