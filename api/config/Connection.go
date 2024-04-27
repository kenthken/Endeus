package config

import (
	"fmt"
	"net/url"

	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	var err error = nil
	var db *gorm.DB = nil
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		EnvConfigs.DBUser,
		EnvConfigs.DBPass,
		EnvConfigs.DBHost,
		EnvConfigs.DBPort,
		EnvConfigs.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		}})

	if err != nil {
		log.Fatal("Cannot connected database ", err)
		return nil
	}

	mySqlDB, err := db.DB()
	err = mySqlDB.Ping()

	if err != nil {
		log.Fatal("Request Timeout ", err)
		return nil
	}

	mySqlDB.SetMaxIdleConns(10)
	mySqlDB.SetConnMaxIdleTime(time.Minute * 3)
	mySqlDB.SetMaxOpenConns(10)
	mySqlDB.SetConnMaxLifetime(time.Minute * 3)

	log.Info("Connected Database " + EnvConfigs.DBDriver + " -- running in -- " + EnvConfigs.ClientOrigin)

	return db
}
