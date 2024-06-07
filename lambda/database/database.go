package database

import (
	"whisper-lambda/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "convoGroups"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func (d DynamoDBClient) CreateGroup(group types.Group) error {
	groupItem := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"groupName": {
				S: aws.String(group.GroupName),
			},
			"createdBy": {
				S: aws.String(group.CreatedBy),
			},
			"groupID": {
				S: aws.String(group.GroupName),
			},
		},
	}

	_, err := d.databaseStore.PutItem(groupItem)

	if err != nil {
		return err
	}
	return nil
}

func NewDynamoDBInstance() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	dbInstance := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: dbInstance,
	}

}
