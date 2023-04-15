package database

import (
	"be-ifid/config"
	"be-ifid/utils"
	"context"
	"fmt"

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

	db.SetConnMaxIdleTime(conf.DB.ConnectionIdle)
	db.SetConnMaxLifetime(conf.DB.ConnectionLifetime)
	db.SetMaxIdleConns(conf.DB.MaxIdle)
	db.SetMaxOpenConns(conf.DB.MaxOpen)

	postgresConn = new(sqlx.DB)
	postgresConn = db

	utils.Success("Postgres connected")
}

func PostgresConnection() *sqlx.DB {
	return postgresConn
}

func PostgresIsConnected() bool {
	err := postgresConn.Ping()
	if err != nil {
		utils.SafeError("Postgres fails health check: ", err)
		return false
	}
	return true
}
