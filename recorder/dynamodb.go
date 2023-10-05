package recorder

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/go-spectest/spectest"
)

// NewDynamoDB wraps a DynamoDB client and records all requests and responses.
func NewDynamoDB(cli dynamodbiface.DynamoDBAPI, recorder *spectest.Recorder) dynamodbiface.DynamoDBAPI {
	return &dynamoDBRecorder{
		DynamoDBAPI: cli,
		recorder:    recorder,
	}
}

type dynamoDBRecorder struct {
	dynamodbiface.DynamoDBAPI
	recorder *spectest.Recorder
}

func (a dynamoDBRecorder) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	a.recordInput("QueryInput", input.String())

	output, err := a.DynamoDBAPI.Query(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("QueryOutput", body, err)

	return output, err
}

func (a dynamoDBRecorder) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	a.recordInput("UpdateItemInput", input.String())

	output, err := a.DynamoDBAPI.UpdateItem(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("UpdateItemOutput", body, err)

	return output, err
}

func (a dynamoDBRecorder) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	a.recordInput("PutItemInput", input.String())

	output, err := a.DynamoDBAPI.PutItem(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("PutItemOutput", body, err)

	return output, err
}

func (a dynamoDBRecorder) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	a.recordInput("GetItemInput", input.String())

	output, err := a.DynamoDBAPI.GetItem(input)

	var body string
	if output != nil {
		body = output.String()
	}

	a.recordOutput("GetItemOutput", body, err)

	return output, err
}

func (a dynamoDBRecorder) recordInput(operation, body string) {
	a.recorder.AddMessageRequest(spectest.MessageRequest{
		Source:    spectest.SystemUnderTestDefaultName,
		Target:    "DynamoDB",
		Header:    operation,
		Body:      body,
		Timestamp: time.Now(),
	})
}

func (a dynamoDBRecorder) recordOutput(operation, body string, err error) {
	if err != nil {
		a.recorder.AddMessageResponse(spectest.MessageResponse{
			Source:    "DynamoDB",
			Target:    spectest.SystemUnderTestDefaultName,
			Header:    "Error",
			Body:      fmt.Sprintf("Error: %s", err.Error()),
			Timestamp: time.Now(),
		})
	} else {
		a.recorder.AddMessageResponse(spectest.MessageResponse{
			Source:    "DynamoDB",
			Target:    spectest.SystemUnderTestDefaultName,
			Header:    operation,
			Body:      body,
			Timestamp: time.Now(),
		})
	}
}
