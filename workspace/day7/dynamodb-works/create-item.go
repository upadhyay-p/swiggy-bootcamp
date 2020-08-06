package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	//"types"
)

type Music struct {
	Artist string
	SongTitle string
	AlbumTitle string
	Awards int
}

//insert record into table-->Music2

func main () {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region: aws.String("us-east-1"),
	}))

	db := dynamodb.New(sess)

	music := Music{
		Artist: "Alec",
		SongTitle: "Boy in a bubble",
		AlbumTitle: "Album1",
		Awards: 4,
	}

	musicMap, err := dynamodbattribute.MarshalMap(music)

	if err!=nil {
		panic("Cannot map the values given in music struct..")
	}

	params := &dynamodb.PutItemInput{
		TableName:aws.String("Music2"),
		Item: musicMap,
	}

	resp,err := db.PutItem(params)

	if err!=nil {
		panic("Some error in inserting item.")
	}

	fmt.Println("Record inserted.")
	fmt.Println(resp)

}