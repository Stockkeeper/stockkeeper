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

func GetBlob(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if repository == nil {
			sendResponseRepositoryNameNotFound(
				c,
				fmt.Errorf(
					"no results for a repository with the name \"%v\"",
					repositoryName,
				),
			)
			return
		}

		digest := c.Param("digest")
		if err := registry.ValidateBlobDigest(digest); err != nil {
			sendResponseBlobDigestInvalid(c, err)
			return
		}

		blob, err := registrySrv.GetBlobByRepositoryAndDigest(repository, digest)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get blob by repository and digest",
				err,
			)
			return
		}
		if blob == nil {
			sendResponseBlobDigestNotFound(
				c,
				fmt.Errorf(
					"no results for an image manifest with the reference \"%v\"",
					manifestRef,
				),
			)
			return
		}
	}
}

func GetImageManifest(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if repository == nil {
			sendResponseRepositoryNameNotFound(
				c,
				fmt.Errorf(
					"no results for a repository with the name \"%v\"",
					repositoryName,
				),
			)
			return
		}

		manifestRef := c.Param("manifestRef")
		if err := registry.ValidateImageManifestRef(manifestRef); err != nil {
			sendResponseImageManifestRefInvalid(c, err)
			return
		}

		manifest, err := registrySrv.GetImageManifestByRepositoryAndRef(repository, manifestRef)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if manifest == nil {
			sendResponseImageManifestRefNotFound(
				c,
				fmt.Errorf(
					"no results for an image manifest with the reference \"%v\"",
					manifestRef,
				),
			)
			return
		}

		blob, err := registrySrv.GetBlobByID(manifest.BlobID)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get a blob by ID.",
				err,
			)
			return
		}
		if blob == nil {
			sendResponseBlobNotFound(
				c,
				fmt.Errorf(
					"no results for a blob with the ID \"%v\"",
					manifest.BlobID,
				),
			)
			return
		}

		if err := registrySrv.WriteBlob(blob, c.Writer); err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to write blob to the HTTP response.",
				err,
			)
			return
		}

		c.Status(http.StatusOK)
	}
}

func OpenBlobUploadSession(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if repository == nil {
			sendResponseRepositoryNameNotFound(
				c,
				fmt.Errorf(
					"no results for a repository with the name \"%v\"",
					repositoryName,
				),
			)
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

		digest := c.Query("digest")
		if digest == "" {
			c.Status(http.StatusAccepted)
			c.Header("location", blobUploadSession.Location)
			return
		}

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
	}
}

func UploadChunk(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		if err := registry.ValidateRepositoryName(repositoryName); err != nil {
			sendResponseRepositoryNameInvalid(c, err)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if repository == nil {
			sendResponseRepositoryNameNotFound(
				c,
				fmt.Errorf(
					"no results for a repository with the name \"%v\"",
					repositoryName,
				),
			)
			return
		}

		sessionID, err := uuid.Parse(c.Param("sessionID"))
		if err != nil {
			sendResponseBlobUploadSessionIDInvalid(c, err)
			return
		}

		blobUploadSession, err := registrySrv.GetBlobUploadSessionByRepositoryAndID(repository, sessionID)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get blob upload session by repository and ID.",
				err,
			)
			return
		}
		if blobUploadSession == nil {
			sendResponseBlobUploadSessionNotFound(
				c,
				fmt.Errorf(
					"no results for a blob upload session in the repository \"%v\" and the ID \"%v\"",
					repository.Name,
					sessionID,
				),
			)
			return
		}

		size_uint64, err := strconv.ParseUint(c.GetHeader("content-length"), 10, 64)
		if err != nil {
			sendResponseContentLengthHeaderInvalid(c, err)
			return
		}
		size := uint(size_uint64)

		contentRange := c.GetHeader("content-range")
		if contentRange == "" {
			err := fmt.Errorf("missing request header \"content-range\"")
			sendResponseContentRangeHeaderInvalid(c, err)
			return
		}

		rangeStart, rangeEnd, err := parseContentRange(contentRange)
		if err != nil {
			sendResponseContentRangeHeaderInvalid(c, err)
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
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get repository by name.",
				err,
			)
			return
		}
		if repository == nil {
			sendResponseRepositoryNameNotFound(
				c,
				fmt.Errorf(
					"no results for a repository with the name \"%v\"",
					repositoryName,
				),
			)
			return
		}

		sessionID, err := uuid.Parse(c.Param("sessionID"))
		if err != nil {
			sendResponseBlobUploadSessionIDInvalid(c, err)
			return
		}

		blobUploadSession, err := registrySrv.GetBlobUploadSessionByRepositoryAndID(repository, sessionID)
		if err != nil {
			sendResponseError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_SERVER_ERROR",
				"Error occurred when trying to get blob upload session by repository and ID.",
				err,
			)
			return
		}
		if blobUploadSession == nil {
			sendResponseBlobUploadSessionNotFound(
				c,
				fmt.Errorf(
					"no results for a blob upload session in the repository \"%v\" and the ID \"%v\"",
					repository.Name,
					sessionID,
				),
			)
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
