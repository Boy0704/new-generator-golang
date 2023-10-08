package ginx

import (
	"github.com/gin-gonic/gin"
)

type generalResponse struct {
	RequestID    string      `json:"request_id,omitempty"`
	Message      string      `json:"message,omitempty"`
	Status       int         `json:"status,omitempty"`
	Error        interface{} `json:"error,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	DashboardAll interface{} `json:"dashboard-all-transaction-client-hour,omitempty"`
}

func RespondWithError(ctx *gin.Context, status int, message string, error interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		RequestID: ctx.Value("request_id").(string),
		Message:   message,
		Status:    status,
		Error:     error,
	})
	ctx.Abort()
	return
}

func RespondWithJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		RequestID: ctx.Value("request_id").(string),
		Message:   message,
		Status:    status,
		Data:      data,
	})
	ctx.Abort()
	return
}

func RespondWithCustomJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, data)
	ctx.Abort()
	return
}
