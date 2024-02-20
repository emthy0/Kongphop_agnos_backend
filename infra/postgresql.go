package infra

import "database/sql"

var DB *sql.DB

func InitializeDB(connectionString string) error {
  var err error
  DB, err = sql.Open("postgres", connectionString)
  if err != nil {
      return err
  }

  err = DB.Ping()
  if err != nil {
      return err
  }

  return nil
}

func ProvideDB() *sql.DB {
  return DB
}

