package gofire

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"google.golang.org/api/iterator"
)

type Firestore struct {
	client          *Client
	firestoreClient *firestore.Client
	ctx             context.Context
}

type Clause struct {
	FieldName string
	Operator  string
	Value     interface{}
}

func (f Firestore) CreateDocument(collectionName, documentId string, documentToCreate interface{}) (*firestore.WriteResult, error) {
	return f.firestoreClient.Collection(collectionName).Doc(documentId).Set(f.ctx, documentToCreate)
}

func (f Firestore) GetDocument(collectionName, documentId string) (string, error) {
	establishmentSnap, err := f.firestoreClient.Collection(collectionName).Doc(documentId).Get(f.ctx)
	CheckErr(err)

	establishmentMap := establishmentSnap.Data()
	establishmentJson, err := json.Marshal(establishmentMap)
	CheckErr(err)

	return string(establishmentJson), err
}

func (f Firestore) UpdateDocument(collectionName, documentId string, updateList []map[string]interface{}) (*firestore.WriteResult, error) {
	var allUpdates []firestore.Update = make([]firestore.Update, 0)

	for _, oneUpdateValue := range updateList { // runs through all the fields to update

		for fieldPath, fieldValue := range oneUpdateValue {
			var oneUpdate firestore.Update = firestore.Update{
				Path:  fieldPath,
				Value: fieldValue,
			}

			allUpdates = append(allUpdates, oneUpdate)
		}

	}

	res, err := f.firestoreClient.Collection(collectionName).Doc(documentId).Update(f.ctx, allUpdates)

	return res, err
}

func (f Firestore) GetDocumentList(collectionName string, clauseList []Clause) string {
	var allDocuments []map[string]interface{} = make([]map[string]interface{}, 0)

	collectionRef := f.firestoreClient.Collection(collectionName)

	for _, oneClause := range clauseList {
		collectionRef.Where(oneClause.FieldName, oneClause.Operator, oneClause.Value)
	}

	res := collectionRef.Documents(f.ctx)

	defer res.Stop()
	for {
		establishmentSnap, err := res.Next()
		if err == iterator.Done {
			break
		}
		CheckErr(err)

		establishmentMap := establishmentSnap.Data()

		allDocuments = append(allDocuments, establishmentMap)
	}

	allDocumentsJson, err := json.Marshal(allDocuments)
	CheckErr(err)

	return string(allDocumentsJson)
}

func (f Firestore) DeleteDocument(collectionName, documentId string) (*firestore.WriteResult, error) {
	result, err := f.firestoreClient.Collection(collectionName).Doc(documentId).Delete(f.ctx)

	return result, err
}
