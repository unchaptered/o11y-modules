package factory

import (
	awsdynomowrapper "o11y-modules/sdk/aws_dynomo_wrapper"
	"testing"
)

func TestDynamoDbFactory(t *testing.T) {
	dynamoDBFactory := DynamoDBFactory{}
	dynamoDB := dynamoDBFactory.GetDynamoDB()

	_, isOk := dynamoDB.(*awsdynomowrapper.DynamoDB)
	if !isOk {
		t.Error(
			"DyanmoFactory{}.GetDnyamoDB()가 awsdynamowrapper.DynamoDB를 리턴하지 않습니다.",
		)
	}
}