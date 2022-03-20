package server

import (
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/Stockkeeper/stockkeeper/config"
	"github.com/Stockkeeper/stockkeeper/frontend"
	"github.com/Stockkeeper/stockkeeper/registry"
	"github.com/Stockkeeper/stockkeeper/server/registry_api"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func NewServer(config config.Config, registryService *registry.Service) *http.Server {
	router := gin.Default()

	{
		frontendRouter := router.Group("/")

		frontendFS := frontend.GetFileSystem()
		fs.WalkDir(
			frontendFS,
			".",
			func(path string, d fs.DirEntry, err error) error {
				if d.IsDir() {
					return nil
				}

				frontendRouter.GET("/"+path, func(c *gin.Context) {
					file, err := frontendFS.Open(path)
					if err != nil {
						panic(err)
					}

					data, err := io.ReadAll(file)
					if err != nil {
						panic(err)
					}

					mime := mimetype.Detect(data).String()
					if strings.Contains(mime, "text/plain") {
						if strings.HasSuffix(path, ".css") {
							mime = "text/css"
						} else if strings.HasSuffix(path, ".js") {
							mime = "text/javascript"
						}
					}

					c.Header("Content-Type", mime)
					c.Writer.Write(data)
				})

				if path == "index.html" {
					frontendRouter.GET("/", func(c *gin.Context) {
						file, err := frontendFS.Open(path)
						if err != nil {
							panic(err)
						}
						io.Copy(c.Writer, file)
					})
				}

				return nil
			},
		)
	}

	{
		apiV1Router := router.Group("/api/v1")
		apiV1Router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World!")
		})
	}

	{
		registryAPIRouter := router.Group("/v2")
		registryAPIRouter.POST("/:repository_name/blobs/uploads", registry_api.HandlePostBlobUploadSession(registryService))
	}

	return &http.Server{
		Addr:    listenAddress,
		Handler: router,
	}
}
