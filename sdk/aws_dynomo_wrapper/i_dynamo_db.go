package awsdynomowrapper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type IDynamoDB interface {
	IsExistsTable(tableName string, sessionData session.Session) (*dynamodb.DescribeTableOutput, error)
	CreateTable(sessionData session.Session) (*dynamodb.CreateTableOutput, error)
}