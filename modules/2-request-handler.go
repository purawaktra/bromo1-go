package modules

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo1-go/dto"
	"github.com/purawaktra/bromo1-go/utils"
	"go.mongodb.org/mongo-driver/bson"
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

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "application/bson" {
		utils.Error(errors.New("contentType not match"), "RetrieveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse data to struct and check err
	var requestBody dto.Request
	err = bson.Unmarshal(readData, &requestBody)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("RetrieveProfilePicture", requestBody)
	response, err := rh.ctrl.RetrieveProfilePicture(requestBody)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// parse response to bson data
	result, err := bson.Marshal(response)
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("Content-Type", "application/bson")
	c.Data(http.StatusOK, "", result)
	return
}

func (rh Bromo1RequestHandler) StoreProfilePicture(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "application/bson" {
		utils.Error(errors.New("contentType not match"), "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse data to struct and check err
	var requestBody dto.Request
	err = bson.Unmarshal(readData, &requestBody)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("StoreProfilePicture", requestBody)
	response, err := rh.ctrl.StoreProfilePicture(requestBody)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// parse response to bson data
	result, err := bson.Marshal(response)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("Content-Type", "application/bson")
	c.Data(http.StatusOK, "", result)
	return
}

func (rh Bromo1RequestHandler) RemoveProfilePicture(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "application/bson" {
		utils.Error(errors.New("contentType not match"), "RemoveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse data to struct and check err
	var requestBody dto.Request
	err = bson.Unmarshal(readData, &requestBody)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	utils.Debug("RemoveProfilePicture", requestBody)
	response, err := rh.ctrl.RemoveProfilePicture(requestBody)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// parse response to bson data
	result, err := bson.Marshal(response)
	if err != nil {
		utils.Error(err, "RemoveProfilePicture", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("Content-Type", "application/bson")
	c.Data(http.StatusOK, "", result)
	return
}
