package password

import (
	"authen.agnoshealth.com/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc domain.PasswordService
}

func (h *handler) GetMinStep() gin.HandlerFunc {
	return func (c *gin.Context) {
		var req StrongPasswordStepRequest
		if c.ShouldBind(&req) != nil  {
			c.String(500, "Invalid input")
		}
		pwd := domain.NewPassword(req.InitPassword)
		numStep := pwd.GetMinSteps()
		c.JSON(200, StrongPasswordStepResponse{
			NumSteps:  numStep,
		})
	}
}