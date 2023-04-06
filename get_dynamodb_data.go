package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamodbData(cfg aws.Config, ctx context.Context, tableName string, out interface{}) error {
	svc := dynamodb.NewFromConfig(cfg)
	result, err := svc.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return err
	}

	return attributevalue.UnmarshalListOfMaps(result.Items, out)
}
