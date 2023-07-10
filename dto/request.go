package dto

type Request struct {
	RequestId string `bson:"request_id"`
	Data      any    `bson:"data"`
}

type RequestProfilePicture struct {
	DocumentId string `bson:"document_id"`
	AccountId  string `bson:"account_id"`
	File       []byte `bson:"data"`
}
