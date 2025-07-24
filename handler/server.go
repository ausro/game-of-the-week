package handler

import (
	"fmt"

	"github.com/ausro/game-of-the-week/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	PORT = 8080
)

type Server struct {
	Gin *gin.Engine
	db  db.Database
}

func New(db db.Database) *Server {
	ginEngine := gin.Default()

	corsConf := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}

	ginEngine.Use(cors.New(corsConf))
	ginEngine.MaxMultipartMemory = 8 << 20
	ginEngine.SetTrustedProxies(nil)

	ginEngine.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	return &Server{
		Gin: ginEngine,
		db:  db,
	}
}

func (server *Server) Run() error {
	return server.Gin.Run(fmt.Sprintf(":%d", PORT))
}
