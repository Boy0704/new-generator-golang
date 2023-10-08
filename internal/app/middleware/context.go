package middleware

import (
	"context"
	"net/http"
	"time"

	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/ginx"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Set("request_id", uuid.New().String())

		done := make(chan struct{})
		go func() {
			defer close(done)
			c.Next()
		}()

		// Wait for the handler to complete or the timeout to occur
		select {
		case <-done:
			// Handler completed within the timeout
			return
		case <-ctx.Done():
			// Handler was canceled or timed out
			if ctx.Err() == context.DeadlineExceeded {
				ginx.RespondWithJSON(c, http.StatusRequestTimeout, "request timeout", nil)
				return
			}
			if ctx.Err() == context.Canceled {
				ginx.RespondWithJSON(c, http.StatusServiceUnavailable, "service unavailable", nil)
				return
			}
		}
	}
}
