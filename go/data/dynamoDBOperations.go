package data

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetItemsList(tableName string) []map[string]*dynamodb.AttributeValue{

	svc := CreateDynamoSessionClient();

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName), // Required
		ExpressionAttributeValues:map[string]*dynamodb.AttributeValue{
			":falseFlag": {
				BOOL: aws.Bool(false),
			},
		},
		FilterExpression: aws.String("IsDeleted = :falseFlag"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == "500" {
			// Specific error code handling
		}
		//		return err
	}
	return result.Items
}

func UnmarshalListOfMaps(l []map[string]*dynamodb.AttributeValue, out interface{}) error {
	attrs := make([]*dynamodb.AttributeValue, len(l))

	for i, m := range l {
		attrs[i] = &dynamodb.AttributeValue{M: m}
	}

	return dynamodbattribute.UnmarshalList(attrs, out)
}

func CreateDynamoSessionClient() *dynamodb.DynamoDB {

	/**
	 * Don't hard-code your credentials!
	 * Export the following environment variables instead:
	 *
	 * export AWS_ACCESS_KEY_ID='AKID'
	 * export AWS_SECRET_ACCESS_KEY='SECRET'
	 */

	// Set your region for future requests.
	//aws.Re Config.Region = "us-west-2"

	aws_access_key_id := "AKIAJTWPQBPM4ZCCNHXQ"
	aws_secret_access_key := "GHVOxB0WnnbvWTv9sII53vKTjVogzwyKUDrvgNhb"
	token := ""

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}


	// Create a session to share configuration, and load external configuration.
	sess := session.Must(session.NewSession())

	// Create the service's client with the session.
	svc := dynamodb.New(sess, aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds))

	return svc
	/*
	//result, err := svc.BatchGetItem(params)
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == "500" {
			// Specific error code handling
		}
		//		return err
	}

	log.Println("Tables:")
	for _, table := range result.TableNames {
		log.Println(*table)
	}*/
}
