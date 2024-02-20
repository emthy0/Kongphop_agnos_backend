package log

import "authen.agnoshealth.com/core/log/repo"

type Log interface {
  SaveLogEntry(req, res string, code int) error
}

func NewLog() Log {
  return repo.NewLogRepository()
}