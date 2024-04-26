package postgre

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-jm-core/pkg/config/structures"
	"go-jm-core/pkg/logger"
	"log"
	"strconv"
	"time"
)

func DbURL(dbConfig *structures.DbConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Db,
	)
}

func OpenPostgresConnection(dbConfig *structures.DbConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", DbURL(dbConfig))
	if err != nil {
		log.Fatal(err)
	}

	maxPoolConnections, err := strconv.Atoi(dbConfig.MaxPoolConnections)
	if err == nil {
		db.DB().SetMaxOpenConns(maxPoolConnections)
	}

	maxIdlePoolConnections, err := strconv.Atoi(dbConfig.MaxIdlePoolConnections)
	if err == nil {
		db.DB().SetMaxIdleConns(maxIdlePoolConnections)
	}

	connectionTimeoutSeconds, err := strconv.Atoi(dbConfig.ConnectionTimeoutSeconds)
	if err == nil {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(connectionTimeoutSeconds))
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	stats := db.DB().Stats()
	logger.Info(fmt.Sprintf("{DUSER:%v, Idle:%v, OpenConnections:%v, InUse:%v, WaitCount:%v, WaitDuration:%v, MaxIdleClosed:%v, MaxLifetimeClosed:%v}",
		dbConfig.Username, stats.Idle, stats.OpenConnections, stats.InUse, stats.WaitCount, stats.WaitDuration, stats.MaxIdleClosed, stats.MaxLifetimeClosed))

	return db, nil
}

func CloseDbConnection(db *gorm.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
