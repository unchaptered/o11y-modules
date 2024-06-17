package factory

import (
	awsdynomowrapper "o11y-modules/sdk/aws_dynomo_wrapper"
)

type DynamoDBFactory struct{}

func (a *DynamoDBFactory) GetDynamoDB() awsdynomowrapper.IDynamoDB {
	return &awsdynomowrapper.DynamoDB{}
}