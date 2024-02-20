package log

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func RequestResponseLogger() gin.HandlerFunc {
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
      err = saveLogEntry(requestString, responseBody, statusCode)
      if err != nil {
          log.Println("Error saving log entry:", err)
      }
  }
}

