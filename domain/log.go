package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type (
	Log struct {
		Request string
    Response string
    Code int
	}
	LogRepository interface {
		SaveLog(ctx context.Context,log *Log) ( error)
	}

	LogService interface {
		WriteLog(ctx context.Context,log *Log) ( error)
	}

	LogMiddleware interface {
		LogReqRes() gin.HandlerFunc
	}
)