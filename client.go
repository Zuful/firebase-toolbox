package go_fire

import (
	firebase "firebase.google.com/go"
)

var app *firebase.App
var firebaseErr error

type Client struct {
	credentialFilePath string
	bucketName         string
}

func NewClient(bucketName string) *Client {
	return &Client{}
}

func (client *Client) NewFirestore() *Firestore {
	return &Firestore{}
}

func (client *Client) NewStorage() *Storage {
	return &Storage{}
}
