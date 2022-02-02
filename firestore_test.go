package gofire_test

import (
	"fmt"
	"github.com/Zuful/gofire"
	"testing"
)

var (
	client                = gofire.NewClient("", "")
	firestore             = client.NewFirestore()
	collectionName        = "images-infos"
	documentId     string = "1"
)

func TestCreateDocument(t *testing.T) {
	var imgInfos = struct {
		Name     string `json:"name"`
		Url      string `json:"url"`
		Category string `json:"category"`
	}{
		Name:     "Test",
		Url:      "https://www.test.com/",
		Category: "teenager",
	}

	_, err := firestore.CreateDocument(collectionName, documentId, imgInfos)

	if err != nil {
		t.Error(err)
	}
}

func TestGetDocument(t *testing.T) {
	document, err := firestore.GetDocument(collectionName, documentId)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(document)
	}
}

func TestUpdateDocument(t *testing.T) {
	var updateList []map[string]interface{} = []map[string]interface{}{
		{"Name": "updated"},
		{"Url": "https://updated.com/"},
	}

	_, err := firestore.UpdateDocument(collectionName, documentId, updateList)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDocument(t *testing.T) {
	_, err := firestore.DeleteDocument(collectionName, documentId)

	if err != nil {
		t.Error(err)
	}
}
