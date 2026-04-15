package engine

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/alwindoss/duka/ui"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Addr string
}

func Start(cfg *Config) error {
	r := gin.Default()

	// Create the sub-filesystem for the internal 'dist' folder
	staticFS, err := fs.Sub(ui.FS, "dist")
	if err != nil {
		log.Fatal("Failed to create sub-FS:", err)
	}

	// API Routes should always be about the r.NoRoute
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "alive"})
	})

	// The Catch-All Handler to render the UI
	r.NoRoute(func(c *gin.Context) {
		path := strings.TrimPrefix(c.Request.URL.Path, "/")

		// DIAGNOSTIC: Check if file exists in the embedded FS
		file, err := staticFS.Open(path)
		if err == nil {
			file.Close()
			log.Printf("Serving file: %s", path)
			http.FileServer(http.FS(staticFS)).ServeHTTP(c.Writer, c.Request)
			return
		}

		// If not found, serve index.html (SPA mode)
		log.Printf("File not found: %s. Falling back to index.html", path)
		index, err := fs.ReadFile(staticFS, "index.html")
		if err != nil {
			c.String(404, "index.html not found in embedded FS")
			return
		}
		c.Data(200, "text/html; charset=utf-8", index)
	})

	err = r.Run(cfg.Addr)
	if err != nil {
		return err
	}
	return nil
}
