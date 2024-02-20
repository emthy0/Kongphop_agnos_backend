package main

import (
	"database/sql"
	"net/http"

	"authen.agnoshealth.com/pkg/log"
	"authen.agnoshealth.com/pkg/password"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := sql.Open("postgres","postgresql://pwdapi@logdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.New()
	logMiddleware := log.Wire(db)
	passwordHandler := password.Wire()
	r.Use(logMiddleware.LogReqRes())
	r.POST("/api/strong_password_steps", passwordHandler.GetMinStep())
	http.ListenAndServe(":8000", nil)
}
