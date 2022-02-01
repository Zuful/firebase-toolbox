package gofire

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App
var firebaseErr error

type Client struct {
	credentialFilePath string
	bucketName         string
	app                *firebase.App
}

func NewClient(credentialFilePath, bucketName string) *Client {
	opt := option.WithCredentialsFile(credentialFilePath)
	app, firebaseErr = firebase.NewApp(context.Background(), nil, opt)
	CheckErr(firebaseErr)

	return &Client{
		credentialFilePath: credentialFilePath,
		bucketName:         bucketName,
		app:                app,
	}
}

func (c *Client) NewFirestore() *Firestore {
	//Firestore
	var ctx context.Context = context.Background()
	firestoreClient, err := app.Firestore(ctx)
	CheckErr(err)

	return &Firestore{
		client:          c,
		firestoreClient: firestoreClient,
		ctx:             ctx,
	}
}

func (c *Client) NewStorage() *Storage {
	return &Storage{
		client: c,
		ctx:    context.Background(),
	}
}
