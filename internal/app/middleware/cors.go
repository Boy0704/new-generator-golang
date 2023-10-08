package middleware

import (
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/configuration"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(cfg *configuration.ConfigApp) gin.HandlerFunc {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Cors.AllowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	return cors.New(corsConfig)

}
