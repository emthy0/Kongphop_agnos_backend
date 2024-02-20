package password

import (
	"sync"

	"authen.agnoshealth.com/domain"
	"github.com/google/wire"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

)

func ProvideHandler(svc domain.PasswordService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			svc: svc,
		}
	})

  

	return hdl
}

func ProvideService() *service {
	svcOnce.Do(func() {
		svc = &service{
		}
	})

	return svc
}

var ProviderSet wire.ProviderSet = wire.NewSet(
  ProvideHandler,
  ProvideService,
)
