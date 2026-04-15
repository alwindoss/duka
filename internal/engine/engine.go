package engine

import (
	"net/http"

	"github.com/alwindoss/duka/ui"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Addr string
}

func Start(cfg *Config) error {
	r := gin.Default()

	// 1. Get the embedded filesystem
	staticFS := ui.GetFileSystem()

	// API Routes should always be about the r.NoRoute
	// Grouping the v1 api
	{
		v1 := r.Group("/duka/api/v1")
		v1.POST("/login", func(ctx *gin.Context) {})
		v1.POST("/logout", func(ctx *gin.Context) {})
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "alive"})
		})
		productsRouter := v1.Group("/products")
		productsRouter.GET("/", func(ctx *gin.Context) {
			ctx.JSONP(200, gin.H{"1": "Product 1"})
		})
	}

	// 2. Serve the static files
	// This handles JS, CSS, and images automatically
	r.StaticFS("/ui", http.FS(staticFS))

	// 3. Optional: Redirect root to /ui/index.html
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui/")
	})

	// 4. Handle Flutter's "Deep Linking" (SPAs)
	// If a user refreshes on /ui/settings, Gin will 404.
	// We should serve index.html for unknown routes under /ui.
	r.NoRoute(func(c *gin.Context) {
		// If the request is for the UI path, serve index.html
		c.FileFromFS("/", http.FS(staticFS))
	})

	err := r.Run(cfg.Addr)
	if err != nil {
		return err
	}
	return nil
}
