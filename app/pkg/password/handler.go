package password

import (
	"fmt"
	"net/http"

	"authen.agnoshealth.com/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc domain.PasswordService
}

func (h *handler) GetMinStep() gin.HandlerFunc {
	return func (c *gin.Context) {
		req := new(StrongPasswordStepRequest)
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
		fmt.Println(req)
		pwd := domain.NewPassword(req.InitPassword)
		numStep := pwd.GetMinSteps()
		c.JSON(200, StrongPasswordStepResponse{
			NumSteps:  numStep,
		})
	}
}