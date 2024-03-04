package database

import (
	"mysql-agent/config"
	"mysql-agent/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Connect(config *config.Config, logger *logrus.Logger) *sqlx.DB {
	db, err := sqlx.Connect("mysql", config.MySQLDSN)
	if err != nil {
		logger.Fatalf("Could not connect to MySQL: %v", err)
	}
	return db
}

func GetSlaveStatus(db *sqlx.DB) (*model.SlaveStatus, error) {
	var status model.SlaveStatus
	err := db.Get(&status, "SHOW SLAVE STATUS")
	return &status, err
}
