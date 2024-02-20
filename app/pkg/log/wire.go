//go:build wireinject
// +build wireinject

package log

import (
	"database/sql"

	"github.com/google/wire"
)

func Wire(db *sql.DB) *middleware {
	panic(wire.Build(ProviderSet))
}