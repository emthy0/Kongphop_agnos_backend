package password

import (
	"context"

	"authen.agnoshealth.com/domain"
)

type service struct {


}

func (s *service) GetMinStep(ctx context.Context, pwd string) (int, error) {
	passwd := domain.NewPassword(pwd)
  return passwd.GetMinSteps(), nil
}