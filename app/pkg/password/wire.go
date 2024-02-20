//go:build wireinject
// +build wireinject

package password

import (
	"github.com/google/wire"
)

func Wire() *handler {
	panic(wire.Build(ProviderSet))
}