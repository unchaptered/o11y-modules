package awsdynomowrapper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDB struct {
}
func (d DynamoDB) IsExistsTable(tableName string, sessionData session.Session) (*dynamodb.DescribeTableOutput, error) {
	svc := dynamodb.New(&sessionData)

	describeResponse, describeError := svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if describeError == nil {
		return nil, describeError
	}

	return describeResponse, nil
}

func (d DynamoDB) CreateTable(sessionData session.Session) (*dynamodb.CreateTableOutput, error) {
	tableName := "example.table"
	_, describeError := d.IsExistsTable(tableName,
														sessionData)
	if describeError == nil {
		return nil, nil
	}

	svc := dynamodb.New(&sessionData)
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("TraceID"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("ThreadID"),
				KeyType: aws.String("RANGE"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("TraceID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ThreadID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("SpanID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ParentID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ServiceName"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("SpanName"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Timestamp"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Duration"),
				AttributeType: aws.String("N"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("SpanIDIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("SpanID"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("ParentIDIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("ParentID"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("ServiceNameIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("ServiceName"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("SpanNameIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("SpanName"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("TimestampIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Timestamp"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("TraceID"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
			{
				IndexName: aws.String("DurationIndex"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Duration"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("TraceID"),
						KeyType:       aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}
	
	response, err := svc.CreateTable(input)
	if err != nil {
		return nil, err
	}

	return response, nil
}