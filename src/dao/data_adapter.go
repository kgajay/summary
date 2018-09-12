package dao

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"logger"
	"fmt"
	"config"
)

var (
	testdb *gorm.DB
)

func Init(env string) {
	var err error
	conf := config.GetConfig()

	if env == "test" {
		testdb, err = gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		// Connect to Postgres database
		testdb, err = connectToDB(conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port,
			conf.Db.Name, conf.Db.Dialect, conf.Db.SslMode)
		if err != nil {
			logger.Log.Fatal("Connection to Database failed with error: " + err.Error())
		}
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
	// Disable table name's pluralization
	db.SingularTable(true)
	// Enable Logger, show detailed log
	db.LogMode(true)
	return db, nil
}
