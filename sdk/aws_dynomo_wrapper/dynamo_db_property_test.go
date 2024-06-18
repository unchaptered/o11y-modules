package awsdynomowrapper

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)


func TestDynamoDBProeprties(t *testing.T) {

	var dynamo DynamoDB
	
	// Check IsExistsTable Properties
	if _, ok := interface{}(dynamo).(interface {
		IsExistsTable(string, session.Session) (*dynamodb.DescribeTableOutput, error) 
	}); !ok {
		t.Errorf("DynamoDB does not have method IsExistsTable")
	}

	// Check CreateTable Properties
	if _, ok := interface{}(dynamo).(interface {
		CreateTable(session.Session) (*dynamodb.CreateTableOutput, error)
	}); !ok {
		t.Errorf("DynamoDB does not have method CreateTable")
	}
	
}