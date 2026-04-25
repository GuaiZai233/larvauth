package router

import (
	"github.com/GuaiZai233/larvauth/internal/auth"
	"github.com/GuaiZai233/larvauth/internal/profile"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// lightweight health endpoint for container HEALTHCHECKs
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	public := r.Group("/api")
	{
		public.POST("/login", auth.Login)
		public.POST("/reg", auth.Register)
		public.GET("/me", profile.GetProfile)
	}
}
