package gofire

import (
	firebase "firebase.google.com/go"
)

var app *firebase.App
var firebaseErr error

type Client struct {
	credentialFilePath string
	bucketName         string
}

func NewClient(credentialFilePath, bucketName string) *Client {
	return &Client{
		credentialFilePath: credentialFilePath,
		bucketName:         bucketName,
	}
}

func (client *Client) NewFirestore() *Firestore {
	return &Firestore{
		credentialFilePath: client.credentialFilePath,
		bucketName:         client.bucketName,
	}
}

func (client *Client) NewStorage() *Storage {
	return &Storage{
		credentialFilePath: client.credentialFilePath,
		bucketName:         client.bucketName,
	}
}
