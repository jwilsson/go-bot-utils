package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetDynamodbData(session *session.Session, tableName string, out interface{}) error {
	svc := dynamodb.New(session)
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return err
	}

	return dynamodbattribute.UnmarshalListOfMaps(result.Items, &out)
}
