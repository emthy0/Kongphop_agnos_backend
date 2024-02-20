package repo

import (
	"database/sql"
	"time"
)

type LogEntry struct {
  ID        int
  Request   string
  Response  string
  CreatedAt time.Time
}

type LogRepository struct {
  db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
  return &LogRepository{db}
}

func(repo *LogRepository) SaveLogEntry(request, response string, statusCode int) error {
  query := `INSERT INTO logs(request, responsecode, response, created_at) VALUES($1, $2, $3, $4)`
  _, err := repo.db.Exec(query, request, statusCode, response, time.Now())
  if err != nil {
      return err
  }
  return nil
}