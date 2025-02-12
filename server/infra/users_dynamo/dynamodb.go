package users_dynamo

import (
	"context"
	"log"
	"os"

	"github.com/MauriPinoRicci/example-api-go/server/domain/users"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type repo struct {
	client *dynamodb.Client
}

var _ users.Repository = (*repo)(nil) // implement interface

func New() *repo {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)),
	)
	if err != nil {
		log.Fatalf("Error al cargar configuración: %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return &repo{
		client: svc,
	}
}

func (s *repo) Save(ctx context.Context, entity *users.User) error {

	msg := BuildUserMsg(entity)

	item, err := attributevalue.MarshalMap(msg)
	if err != nil {
		return err
	}

	table := "Users"
	_, err =  s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *repo) GetByID(ctx context.Context, id string) (*users.User, error) {
	return nil, nil
}

func (s *repo) Delete(ctx context.Context, id string) error {
	return nil
}
