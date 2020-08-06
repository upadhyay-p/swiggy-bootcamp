package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Music1 struct {
	Artist string
	SongTitle string
	AlbumTitle string
	Awards int
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("1"),
	}))

	db := dynamodb.New(sess)

	artist := "Alec"
	songTitle := "Boy in a bubble"

	params := &dynamodb.GetItemInput{
		TableName: aws.String("Music2"),
		Key: map[string]*dynamodb.AttributeValue{
			"Artist":{
				S:aws.String(artist),
		},
		"SongTitle":{
			S: aws.String(songTitle),
		},
		},
	}

	resp,err := db.GetItem(params)
	if err!=nil {
		panic("Sorry item not found... ")
	}

	fmt.Println(resp.Item)

	var music = Music1{}
	dynamodbattribute.UnmarshalMap(resp.Item, music)

	fmt.Printf("Unmarshalled Item %+v:",music)
}