package dto

type Request struct {
	RequestId  string
	DocumentId string
	AccountId  string
}

type Header struct {
	RequestId  string `header:"request_id" binding:"required"`
	DocumentId string `header:"document_id" binding:"required"`
}

type HeaderStorePhotoProfile struct {
	RequestId string `header:"request_id"`
	AccountId string `header:"account_id"`
}

type RequestProfilePicture struct {
	DocumentId string `bson:"document_id"`
	AccountId  string `bson:"account_id"`
	File       []byte `bson:"data"`
}
