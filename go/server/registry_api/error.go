package registry_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponseBody struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func sendResponseError(c *gin.Context, httpStatus int, code string, message string, err error) {
	c.JSON(
		httpStatus,
		ErrorResponseBody{[]Error{
			{
				Code:    code,
				Message: message,
				Detail:  err.Error(),
			},
		}},
	)
}

func sendResponseBlobUploadSessionIDInvalid(c *gin.Context, err error) {
	sendResponseError(
		c,
		http.StatusBadRequest,
		"BLOB_UPLOAD_INVALID",
		"The blob upload session ID is invalid.",
		err,
	)
}

func sendResponseBlobUploadSessionNotFound(c *gin.Context, err error) {
	sendResponseError(
		c,
		http.StatusNotFound,
		"BLOB_UPLOAD_UNKNOWN",
		"The blob upload session could not be found.",
		err,
	)
}

func sendResponseRepositoryNameInvalid(c *gin.Context, err error) {
	sendResponseError(
		c,
		http.StatusBadRequest,
		"NAME_INVALID",
		"The repository name is invalid.",
		err,
	)
}

func sendResponseRepositoryNameNotFound(c *gin.Context, err error) {
	sendResponseError(
		c,
		http.StatusNotFound,
		"NAME_UNKNOWN",
		"The repository name could not be found.",
		err,
	)
}

func sendResponseContentLengthHeaderInvalid(c *gin.Context, err error) {
	sendResponseError(
		c,
		http.StatusBadRequest,
		"CONTENT_LENGTH_INVALID",
		"The HTTP request header \"content-length\" is invalid.",
		err,
	)
}
