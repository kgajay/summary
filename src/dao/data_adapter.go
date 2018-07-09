package dao

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"logger"
	"fmt"
	"config"
)

var (
	testdb *gorm.DB
)

func Init() {
	var err error
	conf := config.GetConfig()

	// Connect to Postgres database
	testdb, err = connectToDB(conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port,
		conf.Db.Name, conf.Db.Dialect, conf.Db.SslMode)
	if err != nil {
		logger.Log.Fatal("Connection to Database failed with error: " + err.Error())
	}

}

// GetDb for getting DB database
func GetDb() *gorm.DB {
	return testdb
}

func connectToDB(user, password, host, port, dbName, dbType, sslMode string) (db *gorm.DB, err error) {
	logger.Log.Info("Connecting to database")
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, dbName, sslMode)
	db, err = gorm.Open(dbType, connectionString)
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping() // check the database connectivity
	if err != nil {
		return nil, err
	}
	return db, nil
}
