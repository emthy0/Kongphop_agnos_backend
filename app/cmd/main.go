package main

import (
	"database/sql"
	"fmt"
	"os"

	"authen.agnoshealth.com/pkg/log"
	"authen.agnoshealth.com/pkg/password"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	postgresHostname := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPwd := os.Getenv("POSTGRES_PASSWORD")
	dbName :=  os.Getenv("POSTGRES_DB")
	db, err := sql.Open("postgres",fmt.Sprintf("postgresql://%s:%s@%s/%s",postgresUser, postgresPwd, postgresHostname, dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.New()
	logMiddleware := log.Wire(db)
	passwordHandler := password.Wire()
	r.Use(logMiddleware.LogReqRes())
	r.POST("/api/strong_password_steps", passwordHandler.GetMinStep())
	r.Run(":8000")
}
