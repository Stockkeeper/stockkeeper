package registry_api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Stockkeeper/stockkeeper/go/registry"
	"github.com/gin-gonic/gin"
)

func OpenBlobUploadSession(registrySrv *registry.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		repositoryName := c.Param("repositoryName")
		// TODO: Check if repository name is valid.
		if repositoryName == "" {
			c.JSON(
				http.StatusNotFound,
				ErrorReponseBody{[]Error{
					{
						Code:    "NAME_INVALID",
						Message: fmt.Sprintf("The repository name is invalid: %v", repositoryName),
					},
				}},
			)
			return
		}

		repository, err := registrySrv.GetRepositoryByName(repositoryName)
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				ErrorReponseBody{[]Error{
					{
						Code:    "NAME_UNKNOWN",
						Message: fmt.Sprintf("Could not find repository with name %v", repositoryName),
					},
				}},
			)
			return
		}

		blobUploadSession, err := registrySrv.OpenBlobUploadSession(repository)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				ErrorReponseBody{[]Error{
					{
						Code:    "INTERNAL_SERVER_ERROR",
						Message: "Failed to opan a new blob upload session.",
					},
				}},
			)
			return
		}

		if digest := c.Query("digest"); digest != "" {
			size, err := strconv.ParseUint(c.GetHeader("content-length"), 10, 64)
			if err != nil {
				panic(err)
			}

			_, err = registrySrv.AppendChunk(blobUploadSession, c.Request.Body, uint(size), 0, 0)
			if err != nil {
				panic(nil)
			}

			blob, err := registrySrv.CloseBlobUploadSession(blobUploadSession)
			if err != nil {
				panic(nil)
			}

			c.Status(http.StatusCreated)
			c.Header("location", blob.Location)
			return
		}
	}
}
