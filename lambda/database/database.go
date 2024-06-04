package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type IDatabase interface{}

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBInstance() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	dbInstance := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: dbInstance,
	}

}
