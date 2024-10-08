package artifact

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *Database

type Database struct {
	*gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := Config.GetString(Config.RelationDBConfig+".Username") + ":" + Config.GetString(Config.RelationDBConfig+".Password") + "@tcp(" + Config.GetString(Config.RelationDBConfig+".Host") + ":" + Config.GetString(Config.RelationDBConfig+".Port") + ")/" + Config.GetString(Config.RelationDBConfig+".Database") + "?charset=utf8mb4&parseTime=True&loc=Local"
	connection := Config.GetString(Config.RelationDBConfig + ".Connection")
	config := &gorm.Config{}

	if Config.GetString(Config.RelationDBConfig+".DbLogEnabled") == "true" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // Don't include params in the SQL log
				Colorful:                  true,        // Disable color
			},
		)
		config.Logger = newLogger
	}

	err := error(nil)
	var db *gorm.DB

	switch connection {
	case "mysql":
		dsn = Config.GetString(Config.RelationDBConfig+".Username") + ":" + Config.GetString(Config.RelationDBConfig+".Password") + "@tcp(" + Config.GetString(Config.RelationDBConfig+".Host") + ":" + Config.GetString(Config.RelationDBConfig+".Port") + ")/" + Config.GetString(Config.RelationDBConfig+".Database") + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), config)
	case "postgres":
		dsn = "host=" + Config.GetString(Config.RelationDBConfig+".Host") + " user=" + Config.GetString(Config.RelationDBConfig+".Username") + " password=" + Config.GetString(Config.RelationDBConfig+".Password") + " dbname=" + Config.GetString(Config.RelationDBConfig+".Database") + " port=" + Config.GetString(Config.RelationDBConfig+".Port") + " sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), config)
	case "sqlite":
		dsn = Config.GetString(Config.RelationDBConfig + ".Database")
		db, err = gorm.Open(sqlite.Open(dsn), config)
	default:
		err = errors.New("db connection not supported")

	}

	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}
