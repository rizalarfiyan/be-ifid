package database

import (
	"be-ifid/config"
	"be-ifid/utils"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var postgresConn *sqlx.DB

func PostgresInit() {
	utils.Info("Connect postgres server...")
	conf := config.Get()
	ctx := context.Background()
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Name)
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		utils.Error("Postgres connection problem: ", err)
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(100)
	db.SetConnMaxIdleTime(2 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	postgresConn = new(sqlx.DB)
	postgresConn = db

	utils.Success("Postgres connected")
}

func PostgresGet() *sqlx.DB {
	return postgresConn
}
