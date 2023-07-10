package modules

import (
	"context"
	"fmt"
	"github.com/purawaktra/bromo1-go/entities"
	"github.com/purawaktra/bromo1-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bromo1Repo struct {
	db *mongo.Client
}

func CreateBromo1Repo(db *mongo.Client) Bromo1Repo {
	return Bromo1Repo{
		db: db,
	}
}

type Bromo1RepoInterface interface {
	RetrieveProfilePictureByDocumentId(data entities.ProfilePicture) (entities.ProfilePicture, error)
	StoreProfilePicture(data entities.ProfilePicture) (entities.ProfilePicture, error)
	RemoveProfilePictureByDocumentId(data entities.ProfilePicture) error
}

func (br Bromo1Repo) RetrieveProfilePictureByDocumentId(data entities.ProfilePicture) (entities.ProfilePicture, error) {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionProfilePicture)
	filter := bson.M{"_id": data.DocumentId}

	// get document from mongo db
	doc := coll.FindOne(context.TODO(), filter)

	// check if err
	err := doc.Err()
	if err != nil {
		utils.Error(err, "RetrieveProfilePictureById", filter)
		return entities.ProfilePicture{}, err
	}

	// parse document to struct
	var result entities.ProfilePicture
	err = doc.Decode(&result)

	// check if err parse document
	if err != nil {
		utils.Error(err, "RetrieveProfilePictureById", "")
		return entities.ProfilePicture{}, err
	}

	// create return
	return result, nil
}

func (br Bromo1Repo) StoreProfilePicture(data entities.ProfilePicture) (entities.ProfilePicture, error) {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionProfilePicture)

	// create bson data format
	parsedData, err := bson.Marshal(data)

	// check if err
	if err != nil {
		utils.Error(err, "StoreProfilePicture", data)
		return entities.ProfilePicture{}, err
	}

	// push data to mongo db
	doc, err := coll.InsertOne(context.TODO(), parsedData)

	// check if err
	if err != nil {
		utils.Error(err, "StoreProfilePicture", data)
		return entities.ProfilePicture{}, err
	}

	// create return
	data.DocumentId = fmt.Sprintf("%v", doc.InsertedID)
	return data, nil
}

func (br Bromo1Repo) RemoveProfilePictureByDocumentId(data entities.ProfilePicture) error {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionProfilePicture)
	filter := bson.M{"_id": data.DocumentId}

	// delete document on mongo db
	result, err := coll.DeleteOne(context.TODO(), filter)

	// check if err
	if err != nil {
		utils.Error(err, "RemoveProfilePictureByDocumentId", filter)
		return err
	}

	// warn if deleted two
	if result.DeletedCount > 1 {
		utils.Warn("RemoveProfilePictureByDocumentId", "deletd more than one document")
	}

	// create return
	return nil
}
