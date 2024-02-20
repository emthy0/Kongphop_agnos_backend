package router

import (
	"authen.agnoshealth.com/controller/password"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

  router := gin.Default()

  api := router.Group("/api")
  api.POST("/strong_password_steps", password.StrongPasswordStepHandler)

  return router
}