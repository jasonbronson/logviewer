package transport

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

//go:embed frontend/dist
var feserver embed.FS

// Router func
func Router(newRelicApp *newrelic.Application) http.Handler {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}
	router := gin.Default()
	router.RedirectTrailingSlash = true

	router.Use(nrgin.Middleware(newRelicApp))
	router.Use(cors.New(corsConfig))

	//Frontend serve vue ui
	router.Use(static.Serve("/", EmbedFolder(feserver, "frontend/dist")))
	router.NoRoute(func(c *gin.Context) {
		fmt.Printf("%s doesn't exist, redirect on /", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// if os.Getenv("ENVIRONMENT") == "load" {
	// 	pprof.Register(router, "debug")
	// }

	router.GET("/", HealthCheck)
	router.GET("/healthz", HealthCheck)

	//Performance verify key on load forge
	// loaderVerification := os.Getenv("LOAD_FORGE")
	// if loaderVerification != "" {
	// 	router.GET(fmt.Sprintf("/%v", loaderVerification), func(g *gin.Context) {
	// 		g.Writer.WriteHeader(http.StatusOK)
	// 		g.Writer.Write([]byte(loaderVerification))
	// 	})
	// }

	return router
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
