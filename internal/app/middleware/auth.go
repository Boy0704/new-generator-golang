package middleware

import (
	"net/http"

	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/configuration"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/ginx"
	"github.com/gin-gonic/gin"
)

type BasicAuth struct {
	Cfg *configuration.ConfigApp
}

func (m *BasicAuth) CustomerServiceBasicAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, password, _ := context.Request.BasicAuth()
		if ok := m.isAuthorizeBasicAuth(username, password); !ok {
			ginx.RespondWithJSON(context, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
		context.Next()
	}
}

func (m *BasicAuth) isAuthorizeBasicAuth(username, password string) bool {
	if m.Cfg.BasicAuth.Username == username && m.Cfg.BasicAuth.Password == password {
		return true
	}
	return false
}
