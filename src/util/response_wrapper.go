package util

import (
	"apigouniv/src/domain"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

// Log3rdParty is a function to create format log contains method, url, and request, response body. It's created after consume 3rd party.
func Log3rdParty(method, url, req, resp interface{}) {
	// Assign detail with parameter above and generate console
	// Request
	logger.WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": req,
	}).Info("Resty Request")
	Logkoe.Info("Resty Request => method =", method,
		"||url =", url,
		"||request =", req)

	// Response
	logger.WithFields(logger.Fields{
		"response": resp,
	}).Info("Resty Response")
	Logkoe.Info("Resty Response ==> response =", resp)
}

// HandleSuccess is a function of success process that send to the front as a response success
func HandleSuccess(c *gin.Context, status int, code string, message string, data map[string]interface{}, detail interface{}, info string) {
	// Assign struct ResponseWrapper with parameter above
	response := domain.ResponseWrapper{
		StatusCode: code,
		Message:    message,
		Result:     data,
	}

	// Write header manually
	c.Header("Content-Type", "application/json")

	// Use JSON as a response and send initialize struct above
	c.JSON(status, response)
	LogSuccess(detail, message)
}

// LogSuccess is a function of success process that only generate log to console
func LogSuccess(detail interface{}, info string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Info(info)
	Logkoe.Info("msg =", info, "detail =", detail)
}

// HandleError is a function of failed process that send to the front as a response error
func HandleError(c *gin.Context, status int, code string, message string, detail interface{}, info string) {
	// Assign struct ResponseWrapper with parameter above
	response := domain.ResponseWrapper{
		StatusCode: code,
		Message:    message,
		Result:     nil,
	}

	// Write header manually
	c.Header("Content-Type", "application/json")

	// Use JSON as a response and send initialize struct above
	c.JSON(status, response)
	LogError(detail, message)
}

// LogError is a function of failed process that only generate log to console
func LogError(detail interface{}, info string) {
	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Error(info)
	Logkoe.Info("msg =", info, "detail =", detail)
}
