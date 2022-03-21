package server

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/Stockkeeper/stockkeeper/go/config"
	"github.com/Stockkeeper/stockkeeper/go/frontend"
	"github.com/Stockkeeper/stockkeeper/go/registry"
	"github.com/Stockkeeper/stockkeeper/go/server/registry_api"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg config.Config, registrySrv *registry.Service) *http.Server {
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
		registryAPIRouter.POST("/:repository_name/blobs/uploads", registry_api.OpenBlobUploadSession(registrySrv))
	}

	return &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort),
		Handler: router,
	}
}
