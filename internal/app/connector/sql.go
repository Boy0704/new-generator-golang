package connector

import (
	"fmt"
	"time"

	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/configuration"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/logx"
	"github.com/jmoiron/sqlx"
)

func InitSqlConnection(cfg *configuration.ConfigApp) *sqlx.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.HostName,
		cfg.Database.Port,
		cfg.Database.DatabaseName)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		logx.GetLogger().Fatal(err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logx.GetLogger().Fatal(err)
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(1600 * time.Second)

	return db
}
