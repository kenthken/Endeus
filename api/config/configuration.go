package config

import (
	logging "log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	Database setupDatabase
}

type setupDatabase struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	DBName       string `mapstructure:"DB_NAME"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       int    `mapstructure:"DB_PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

var EnvConfigs *setupDatabase

func InitEnvConfigs(gen bool) {
	EnvConfigs = SetupConfiguration(gen)
}

func SetupConfiguration(gen bool) (config *setupDatabase) {

	//DB config using viper
	if gen {
		viper.AddConfigPath("../../../.database")
	} else {
		viper.AddConfigPath("./.database")
	}

	// file name
	viper.SetConfigName("db")

	//type file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Fatal(err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return
}

func InitLogger(db *gorm.DB) {
	newLogger := logger.New(
		logging.New(logging.Writer(), "\r\n", logging.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db.Logger = newLogger
}
