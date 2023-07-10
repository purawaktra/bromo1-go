package modules

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/purawaktra/bromo1-go/dto"
	"github.com/purawaktra/bromo1-go/utils"
	"go.mongodb.org/mongo-driver/bson"
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
	RetrieveProfilePicture(req dto.Request) (any, error)
	StoreProfilePicture(req dto.Request) (any, error)
	RemoveProfilePicture(req dto.Request) error
}

func (ctrl Bromo1Controller) RetrieveProfilePicture(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to bson data
	marshaledData, err := bson.Marshal(req.Data)

	// check err in parsing
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct
	var requestData dto.RequestProfilePicture
	err = bson.Unmarshal(marshaledData, &requestData)

	// check err in parsing
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert string to int and check err
	accountId, err := strconv.Atoi(requestData.AccountId)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the document
	doc, err := ctrl.uc.RetrieveProfilePictureById(requestData.DocumentId, accountId)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success RetrieveProfilePicture",
		Data:    doc,
	}

	// create return
	return result, nil
}

func (ctrl Bromo1Controller) StoreProfilePicture(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to bson data and check err
	marshaledData, err := bson.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestProfilePicture
	err = bson.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert string to int and check err
	accountId, err := strconv.Atoi(requestData.AccountId)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the document and check err
	doc, err := ctrl.uc.StoreProfilePicture(accountId, requestData.File)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", requestData)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success StoreProfilePicture",
		Data:    doc,
	}

	// create return
	return result, nil
}

func (ctrl Bromo1Controller) RemoveProfilePicture(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// marshal to bson data and check err
	marshaledData, err := json.Marshal(req.Data)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// unmarshal to struct and check err
	var requestData dto.RequestProfilePicture
	err = json.Unmarshal(marshaledData, &requestData)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// call usecase for the document and check err
	err = ctrl.uc.RemoveProfilePictureByDocumentId(requestData.DocumentId)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", req.Data)

		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success RemoveProfilePicture",
		Data:    nil,
	}

	// create return
	return result, nil
}
