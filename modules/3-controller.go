package modules

import (
	"github.com/google/uuid"
	"github.com/purawaktra/bromo1-go/dto"
	"github.com/purawaktra/bromo1-go/utils"
	"strconv"
	"time"
)

type Bromo1Controller struct {
	uc Bromo1Usecase
}

func CreateBromo1Controller(uc Bromo1Usecase) Bromo1Controller {
	return Bromo1Controller{uc: uc}
}

type Bromo1ControllerInterface interface {
	RetrieveProfilePicture(req dto.Request) (dto.Response, []byte, error)
	StoreProfilePicture(req dto.Request, data []byte) (dto.Response, error)
	RemoveProfilePicture(req dto.Request) error
}

func (ctrl Bromo1Controller) RetrieveProfilePicture(req dto.Request) (dto.Response, []byte, error) {
	// start timer
	start := time.Now()

	// call usecase for the document
	doc, err := ctrl.uc.RetrieveProfilePictureByDocumentId(req.DocumentId)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", req)

		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, nil, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      "Success",
	}

	// create return
	return result, doc.Data, nil
}

func (ctrl Bromo1Controller) StoreProfilePicture(req dto.Request, data []byte) (dto.Response, error) {
	// start timer
	start := time.Now()

	// parse string to int
	accountId, err := strconv.Atoi(req.AccountId)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", accountId)
		return dto.Response{}, err
	}

	// call usecase for the document and check err
	doc, err := ctrl.uc.StoreProfilePicture(accountId, data)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", req)

		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      doc.DocumentId,
	}

	// create return
	return result, nil
}

func (ctrl Bromo1Controller) RemoveProfilePicture(req dto.Request) (dto.Response, error) {
	// start timer
	start := time.Now()

	// call usecase for the document and check err
	err := ctrl.uc.RemoveProfilePictureByDocumentId(req.DocumentId)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", req)

		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      "Success",
	}

	// create return
	return result, nil
}
