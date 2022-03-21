package registry_api

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Stockkeeper/stockkeeper/go/registry"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OpenBlobUploadSession(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseRepositoryNameNotFound(c, err)
			return
		}

		blobUploadSession, err := registrySrv.OpenBlobUploadSession(repository)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to open a new blob upload session.",
				err,
			)
			return
		}

		if digest := c.Query("digest"); digest != "" {
			size_uint64, err := strconv.ParseUint(c.GetHeader("content-length"), 10, 64)
			if err != nil {
				sendResponseContentLengthHeaderInvalid(c, err)
			}
			size := uint(size_uint64)

			_, err = registrySrv.AppendChunk(blobUploadSession, c.Request.Body, size, 0, size)
			if err != nil {
				sendResponseError(
					c,
					http.StatusInternalServerError,
					"INTERNAL_SERVER_ERROR",
					"Failed to upload the blob and/or chunk.",
					err,
				)
			}

			blob, err := registrySrv.CloseBlobUploadSession(blobUploadSession)
			if err != nil {
				sendResponseError(
					c,
					http.StatusInternalServerError,
					"INTERNAL_SERVER_ERROR",
					"Failed to close the blob upload session.",
					err,
				)
			}

			c.Status(http.StatusCreated)
			c.Header("location", blob.Location)
			return
		}

		c.Status(http.StatusAccepted)
		c.Header("location", blobUploadSession.Location)
	}
}

func UploadChunkAndCloseBlobUploadSession(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseRepositoryNameNotFound(c, err)
			return
		}

		sessionID := c.Param("sessionID")
		if _, err := uuid.Parse(sessionID); err != nil {
			sendResponseBlobUploadSessionIDInvalid(c, err)
			return
		}

		blobUploadSession, err := registrySrv.GetBlobUploadSessionByRepositoryAndID(repository, sessionID)
		if err != nil {
			sendResponseBlobUploadSessionNotFound(c, err)
			return
		}

		size_uint64, err := strconv.ParseUint(c.GetHeader("content-length"), 10, 64)
		if err != nil {
			sendResponseContentLengthHeaderInvalid(c, err)
			return
		}
		size := uint(size_uint64)

		var rangeStart uint = 0
		var rangeEnd uint = size

		if contentRange := c.GetHeader("content-range"); contentRange != "" {
			var err error
			rangeStart, rangeEnd, err = parseContentRange(contentRange)
			if err != nil {
				sendResponseError(
					c,
					http.StatusBadRequest,
					"CONTENT_RANGE_INVALID",
					"The HTTP request header \"content-range\" is invalid.",
					err,
				)
				return
			}

			if rangeStart != blobUploadSession.RangeStart {
				sendResponseError(
					c,
					http.StatusBadRequest,
					"CONTENT_RANGE_INVALID",
					"The HTTP request header \"content-range\" is invalid.",
					fmt.Errorf("the start of the content range (%v) does not match the current byte position of the blob upload session (%v)", rangeStart, blobUploadSession.RangeStart),
				)
				return
			}
		}

		_, err = registrySrv.AppendChunk(blobUploadSession, c.Request.Body, size, rangeStart, rangeEnd)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to upload the blob and/or chunk.",
				err,
			)
			return
		}

		blob, err := registrySrv.CloseBlobUploadSession(blobUploadSession)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to close the blob upload session.",
				err,
			)
		}

		c.Status(http.StatusCreated)
		c.Header("location", blob.Location)
		return

		_, err = registrySrv.AppendChunk(blobUploadSession, c.Request.Body, uint(size), blobUploadSession.RangeStart, uint(size))
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to upload the blob and/or chunk.",
				err,
			)
		}

		blob, err := registrySrv.CloseBlobUploadSession(blobUploadSession)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Failed to close the blob upload session.",
				err,
			)
		}

		c.Status(http.StatusOK)
		c.Header("location", blob.Location)
	}
}

func parseContentRange(s string) (uint, uint, error) {
	matched, err := regexp.MatchString("^[0-9]+-[0-9]+$", s)
	if err != nil {
		return 0, 0, err
	}

	if !matched {
		return 0, 0, fmt.Errorf("regex pattern ^[0-9]+-[0-9]+$ does not match the content-range: %v", s)
	}
	parts := strings.Split(s, "-")
	rangeStart, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	rangeEnd, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return uint(rangeStart), uint(rangeEnd), err
}
