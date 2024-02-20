package password

import (
	"authen.agnoshealth.com/core/password"
	"github.com/gin-gonic/gin"
)

func StrongPasswordStepHandler(c *gin.Context) {
  var req StrongPasswordStepRequest
  if c.ShouldBind(&req) != nil  {
    c.String(500, "Invalid input")
  }
  pwd := password.NewPassword(req.InitPassword)
  numStep := pwd.GetMinSteps()
  c.JSON(200, StrongPasswordStepResponse{
    NumSteps:  numStep,
  })
}