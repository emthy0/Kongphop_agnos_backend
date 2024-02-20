package log

import (
	"bytes"
	"context"
	"log"
	"net/http"

	"authen.agnoshealth.com/domain"
	"github.com/gin-gonic/gin"
)

type middleware struct {
	svc domain.LogService
}

func (m *middleware) LogReqRes() gin.HandlerFunc {
  return func(c *gin.Context) {
    var requestBody bytes.Buffer
    _, err := requestBody.ReadFrom(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
        c.Abort()
        return
    }
    requestString := requestBody.String()
    blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
    c.Writer = blw

    // Call the next handler
    c.Next()

    // Read the response body
    statusCode := c.Writer.Status()
    responseBody := blw.body.String()

    // Save the request and response to the database
    logst := domain.Log{
      Request: requestString,
      Response: responseBody,
      Code: statusCode,
    }
    err = m.svc.WriteLog(context.Background(),&logst)
    if err != nil {
        log.Println("Error saving log entry:", err)
    }
}
}