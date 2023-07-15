package modules

import (
	"errors"
	"fmt"
	"github.com/purawaktra/bromo1-go/entities"
	"github.com/purawaktra/bromo1-go/utils"
	"strconv"
)

type Bromo1Usecase struct {
	repo Bromo1Repo
}

func CreateBromo1Usecase(repo Bromo1Repo) Bromo1Usecase {
	return Bromo1Usecase{
		repo: repo,
	}
}

type Bromo1UsecaseInterface interface {
	RetrieveProfilePictureById(documentId string, accountId int) (ProfilePicture, error)
	StoreProfilePicture(accountId int, data any) (ProfilePicture, error)
	RemoveProfilePictureByDocumentId(documentId string) error
}

func (uc Bromo1Usecase) RetrieveProfilePictureByDocumentId(documentId string) (ProfilePicture, error) {
	// convert input to entity
	doc := entities.ProfilePicture{DocumentId: documentId}

	// call repo for the doc
	docs, err := uc.repo.RetrieveProfilePictureByDocumentId(doc)

	// check for error on call repo
	if err != nil {
		utils.Error(errors.New("error call usecase"), "RetrieveProfilePictureByDocumentId", doc)
		return ProfilePicture{}, err
	}

	// convert entity to dto
	result := ProfilePicture{
		DocumentId: docs.DocumentId,
		AccountId:  strconv.Itoa(int(docs.AccountId)),
		Data:       docs.Data,
	}

	// create return
	return result, nil
}

func (uc Bromo1Usecase) StoreProfilePicture(accountId int, data []byte) (ProfilePicture, error) {
	// check input on accountId
	if accountId < 1 {
		utils.Error(errors.New("accountId can not be zero or negative"), "StoreProfilePicture", accountId)
		return ProfilePicture{}, errors.New("accountId can not be zero or negative")
	}

	// check the size of data and check err if data is more than 1 megabyte
	size := len(data)
	if size > 1048576 {
		utils.Error(errors.New("size data exceeding limit"), "StoreProfilePicture", fmt.Sprintf("Size exceeding 1 megabytes, size : %d", size))
		return ProfilePicture{}, errors.New("size data exceeding limit")
	}

	// convert input to entity
	profilePicture := entities.ProfilePicture{AccountId: uint(accountId), Data: data}

	// call repo for the city
	docs, err := uc.repo.StoreProfilePicture(profilePicture)

	// check for error on call usecase
	if err != nil {
		utils.Error(errors.New("error call usecase"), "StoreProfilePicture", profilePicture)
		return ProfilePicture{}, err
	}

	// convert entity to dto
	result := ProfilePicture{
		DocumentId: docs.DocumentId,
		AccountId:  strconv.Itoa(int(docs.AccountId)),
		Data:       docs.Data,
	}

	// create return
	return result, nil
}

func (uc Bromo1Usecase) RemoveProfilePictureByDocumentId(documentId string) error {
	// skip checking a document documentId
	// convert input to entity
	doc := entities.ProfilePicture{DocumentId: documentId}

	// call repo for the doc
	err := uc.repo.RemoveProfilePictureByDocumentId(doc)

	// check for error on call repo
	if err != nil {
		utils.Error(errors.New("error call usecase"), "RemoveProfilePictureByDocumentId", doc)
		return err
	}

	// create return
	return nil
}
