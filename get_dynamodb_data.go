package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetDynamodbData[T any](ctx context.Context, cfg aws.Config, tableName string) ([]T, error) {
	svc := dynamodb.NewFromConfig(cfg)
	result, err := svc.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	items := make([]T, result.Count)
	err = attributevalue.UnmarshalListOfMaps(result.Items, items)

	return items, err
}
