package data

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	//"net/http"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/awserr"
	//"log"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"log"
)

func GetAgencyList() {

	svc := CreateDynamoSessionClient();

	params := &dynamodb.ScanInput{
		TableName: aws.String("Agency"), // Required
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

	log.Println("Tables:")
	for k, v := range result.Items {
		fmt.Printf("key[%s] value[%s]\n", k, v)

	}
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
