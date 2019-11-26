package db

import (
	"fasoo.com/fklagent/util/config"
	"fasoo.com/fklagent/util/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

var DB *sqlx.DB

func init() {
	initDb()
}

func initDb() {
	log.INFO("db init start!!")
	dbDriver := config.Cfg.DBDriver
	dbHost := config.Cfg.DBHost
	dbPort := config.Cfg.DBPort
	dbUser := config.Cfg.DBUser
	dbPass := config.Cfg.DBPass
	dbName := config.Cfg.DBName

	dbConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	var err error
	DB, err = sqlx.Open(dbDriver, dbConn)
	if err != nil {
		log.ERROR(err.Error())
		panic(err.Error())
	}

	go func() {
		for {
			time.Sleep(time.Minute * 5)
			err := DB.Ping()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
}
