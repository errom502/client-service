package handler

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(vh *VerificationHandler, frontendDir string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	api := r.Group("/api/v1/verification")
	{
		api.GET("/status", vh.CheckStatus)
		api.POST("/create", vh.Create)
		api.POST("/retry", vh.Retry)
	}

	// Раздача React SPA — все неизвестные маршруты отдают index.html
	r.Static("/assets", frontendDir+"/assets")
	r.NoRoute(func(c *gin.Context) {
		c.File(frontendDir + "/index.html")
	})
	_ = http.MethodOptions // suppress unused import warning

	return r
}
