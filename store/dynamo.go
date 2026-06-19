package store

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/umangsarda/golink/models"
)

const tableName = "golink-urls"

var dynamoClient *dynamodb.Client

func InitDynamo() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		return fmt.Errorf("unable to load AWS config: %w", err)
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)
	return nil
}

func SaveLink(link models.Link) error {
	item, err := attributevalue.MarshalMap(link)
	if err != nil {
		return fmt.Errorf("failed to marshal link: %w", err)
	}
	_, err = dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	return err
}

func GetLink(code string) (*models.Link, error) {
	result, err := dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"Code": &types.AttributeValueMemberS{Value: code},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	var link models.Link
	err = attributevalue.UnmarshalMap(result.Item, &link)
	return &link, err
}