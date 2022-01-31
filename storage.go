package gofire

import (
	"bytes"
	"context"
	"firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"net/url"
)

type Storage struct {
	credentialFilePath string
	bucketName         string
}

func (storage *Storage) uploadToFirebase(filePath string) (string, error) {
	//create an id
	id := uuid.New()
	fileInput, err := ioutil.ReadFile(filePath)
	CheckErr(err)
	ctx := context.Background()
	opt := option.WithCredentialsFile(storage.credentialFilePath)
	app, firebaseErr = firebase.NewApp(context.Background(), nil, opt)
	CheckErr(firebaseErr)

	client, err := app.Storage(context.Background())
	CheckErr(err)

	bucket, err := client.Bucket(storage.bucketName)
	CheckErr(err)

	object := bucket.Object(filePath)
	writer := object.NewWriter(ctx)

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return "", err
	}
	/*
		if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
			return "", err
		}
	*/

	var baseStorageImagePath string = "https://firebasestorage.googleapis.com/v0/b/" + storage.bucketName + "/o/" + url.QueryEscape(filePath) + "?alt=media&token="
	var storageImagePath string = baseStorageImagePath + id.String()

	return storageImagePath, nil
}
