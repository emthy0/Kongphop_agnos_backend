package log

import (
	"context"

	"authen.agnoshealth.com/domain"
)

type service struct {
  repo domain.LogRepository
}

func (s *service)  WriteLog(ctx context.Context,log *domain.Log) ( error) {
  return s.repo.SaveLog(ctx, log)
}