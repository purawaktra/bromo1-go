package modules

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo1-go/dto"
	"github.com/purawaktra/bromo1-go/functions"
	"github.com/purawaktra/bromo1-go/utils"
	"io"
	"net/http"
)

type Bromo1RequestHandler struct {
	ctrl Bromo1Controller
}

func CreateBromo1RequestHandler(ctrl Bromo1Controller) Bromo1RequestHandler {
	return Bromo1RequestHandler{
		ctrl: ctrl,
	}
}

type Bromo1RequestHandlerInterface interface {
	RetrieveProfilePicture(c *gin.Context)
	StoreProfilePicture(c *gin.Context)
	RemoveProfilePicture(c *gin.Context)
}

func (rh Bromo1RequestHandler) RetrieveProfilePicture(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", []byte("fail to parse header"))
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", []byte("uuid is not valid"))
		return
	}

	// call controller for the profile picture and check err
	request := dto.Request{
		RequestId:  header.RequestId,
		DocumentId: header.DocumentId,
	}
	utils.Debug("RetrieveProfilePicture", request)
	response, data, err := rh.ctrl.RetrieveProfilePicture(request)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "image/*", data)
	return
}

func (rh Bromo1RequestHandler) StoreProfilePicture(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "image/*" {
		utils.Error(errors.New("contentType not match"), "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get header and check for err
	var header dto.HeaderStorePhotoProfile
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	request := dto.Request{
		RequestId: header.RequestId,
		AccountId: header.AccountId,
	}
	utils.Debug("StoreProfilePicture", header)
	response, err := rh.ctrl.StoreProfilePicture(request, readData)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "", []byte(""))
	return
}

func (rh Bromo1RequestHandler) RemoveProfilePicture(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	request := dto.Request{
		RequestId:  header.RequestId,
		DocumentId: header.DocumentId,
	}
	utils.Debug("RemoveProfilePicture", request)
	response, err := rh.ctrl.RemoveProfilePicture(request)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "", []byte(""))
	return
}
