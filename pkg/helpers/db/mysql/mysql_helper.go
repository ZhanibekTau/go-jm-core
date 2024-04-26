package mysql

import (
	"fmt"
	"github.com/ZhanibekTau/go-jm-core/pkg/config/structures"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func DbURL(dbConfig *structures.DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
	)
}

func NewGormSqlDB(dbConfig *structures.DbConfig) (*gorm.DB, error) {
	Database, err := gorm.Open("mysql", DbURL(dbConfig))
	if err != nil {
		fmt.Println("Status:", err)
		return nil, err
	}

	maxPoolConnections, err := strconv.Atoi(dbConfig.MaxPoolConnections)
	if err == nil {
		Database.DB().SetMaxOpenConns(maxPoolConnections)
	}

	maxIdlePoolConnections, err := strconv.Atoi(dbConfig.MaxIdlePoolConnections)
	if err == nil {
		Database.DB().SetMaxIdleConns(maxIdlePoolConnections)
	}

	connectionTimeoutSeconds, err := strconv.Atoi(dbConfig.ConnectionTimeoutSeconds)
	if err == nil {
		Database.DB().SetConnMaxLifetime(time.Second * time.Duration(connectionTimeoutSeconds))
	}

	err = Database.DB().Ping()
	if err != nil {
		return nil, err
	}
	Database.AutoMigrate()

	return Database, nil
}

func CloseDbConnection(db *gorm.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
