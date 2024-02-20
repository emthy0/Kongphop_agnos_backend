package log

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"authen.agnoshealth.com/domain"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type middleware struct {
	svc domain.LogService
}

func (m *middleware) LogReqRes() gin.HandlerFunc {
  return func(c *gin.Context) {
    json := make(map[string]interface{})
    err := c.ShouldBindBodyWith(&json, binding.JSON)
		var req string
		if err != nil && len(json) == 0 {
			req = ""
		}
		req = fmt.Sprintf("%v", json)
    // requestString := requestBody.String()
    blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
    c.Writer = blw

    // Call the next handler
    c.Next()

    // Read the response body
    statusCode := c.Writer.Status()
    responseBody := blw.body.String()

    // Save the request and response to the database
    logst := domain.Log{
      Request: req,
      Response: responseBody,
      Code: statusCode,
    }
    err = m.svc.WriteLog(context.Background(),&logst)
    if err != nil {
        log.Println("Error saving log entry:", err)
    }
}
}