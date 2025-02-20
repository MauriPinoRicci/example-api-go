package users_dynamo

import (
	"context"
	"errors"
	"log"

	"github.com/MauriPinoRicci/example-api-go/users/domain/users"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type repo struct {
	client    *dynamodb.Client
	tableName *string
}

var _ users.Repository = (*repo)(nil) // implement interface

func New() *repo {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatalf("Error al cargar configuración: %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	table := "Users"
	return &repo{
		client:    svc,
		tableName: &table,
	}
}

func (s *repo) Save(ctx context.Context, entity *users.User) error {
	msg := BuildUserMsg(entity)

	item, err := attributevalue.MarshalMap(msg)
	if err != nil {
		return err
	}

	_, err = s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: s.tableName,
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *repo) GetByID(ctx context.Context, id string) (*users.User, error) {
	res, err := s.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: s.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Item == nil {
		return nil, errors.New("user not found")
	}

	var msg UserMsg
	err = attributevalue.UnmarshalMap(res.Item, &msg)
	if err != nil {
		return nil, err
	}

	entity, err := msg.ToUser()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *repo) Delete(ctx context.Context, id string) error {
	_, err := s.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: s.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return err
	}
	
	return nil
}
