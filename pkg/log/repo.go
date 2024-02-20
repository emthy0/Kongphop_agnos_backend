package log

import (
	"context"
	"database/sql"
	"time"

	"authen.agnoshealth.com/domain"
)

type repository struct {
  db *sql.DB
}

func (r *repository) SaveLog(ctx context.Context,log *domain.Log) ( error) {
  query := `INSERT INTO logs(request, responsecode, response, created_at) VALUES($1, $2, $3, $4)`
  _, err := r.db.Exec(query, log.Request, log.Code, log.Response, time.Now())
  if err != nil {
      return err
  }
  return nil
}