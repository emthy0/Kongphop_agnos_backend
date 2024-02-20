package infra

import "github.com/google/wire"

var InfraSet = wire.NewSet(ProvideDB)