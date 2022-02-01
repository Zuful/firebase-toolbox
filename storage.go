package gofire

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/url"
)

type Storage struct {
	client *Client
	ctx    context.Context
}

func (s *Storage) UploadToFirebase(filePath string) (string, error) {
	//create an id
	id := uuid.New()
	fileInput, err := ioutil.ReadFile(filePath)
	CheckErr(err)

	client, err := s.client.app.Storage(context.Background())
	CheckErr(err)

	bucket, err := client.Bucket(s.client.bucketName)
	CheckErr(err)

	object := bucket.Object(filePath)
	writer := object.NewWriter(s.ctx)

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

	var baseStorageImagePath string = "https://firebasestorage.googleapis.com/v0/b/" + s.client.bucketName + "/o/" + url.QueryEscape(filePath) + "?alt=media&token="
	var storageImagePath string = baseStorageImagePath + id.String()

	return storageImagePath, nil
}
