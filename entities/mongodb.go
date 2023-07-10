package entities

type ProfilePicture struct {
	DocumentId string `bson:"-"`
	AccountId  uint   `bson:"account_id"`
	Data       []byte `bson:"data"`
}
