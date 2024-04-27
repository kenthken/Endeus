package generate

import (
	"endeus/api/config"
	"endeus/api/entities"
	"time"

	"fmt"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Migrate() {
	config.InitEnvConfigs()
	logEntry := fmt.Sprintf("Auto Migrating...")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.EnvConfigs.DBUser,
		config.EnvConfigs.DBPass,
		config.EnvConfigs.ClientOrigin,
		config.EnvConfigs.DBName)

	//init logger
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // Set the logger for GORM
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	db.AutoMigrate(
		&entities.Recipe{},
		&entities.Ingredient{},
		&entities.IngredientDetail{},
		&entities.Method{},
		&entities.MethodDetail{},
		&entities.User{},
		&entities.Rating{},
		&entities.Discussion{},
		&entities.DiscussionReply{},
	)

	if db != nil && db.Error != nil {
		fmt.Sprintf("%s %s with error %s", logEntry, "Failed", db.Error)
		panic(err)
	}

	fmt.Sprintf("%s %s", logEntry, "Success")
}
