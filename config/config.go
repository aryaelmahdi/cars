package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DriverName string
	ServerPort int
	DBPort     int
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
	Secret     string
}

type DB struct {
	Database *sqlx.DB
}

func Init() *Config {
	res := new(Config)
	res = loadConfig()

	return res
}

func LoadDB(config *Config) *DB {
	var dsn = fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBHost, config.DBPort, config.DBName)

	db := new(DB)
	db.initDB(config.DriverName, dsn)

	return db
}

func (db *DB) initDB(driverName string, dataSourceName string) {
	newDB, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		logrus.Error("Config : cannot connect to database", err.Error())
	}
	err = newDB.Ping()
	if err != nil {
		logrus.Error("Config : connection error", err.Error())
	}

	db.Database = newDB
}

func loadConfig() *Config {
	res := new(Config)
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : cannot load configuration,", err.Error())
		return nil
	}

	if val, found := os.LookupEnv("DriverName"); found {
		res.DriverName = val
	}

	if val, found := os.LookupEnv("ServerPort"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Server Port not found,", err.Error())
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPort"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("DB Port not found,", err.Error())
		}
		res.DBPort = port
	}

	if val, found := os.LookupEnv("DBHost"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DBUsername"); found {
		res.DBUsername = val
	}

	if val, found := os.LookupEnv("DBPassword"); found {
		res.DBPassword = val
	}

	if val, found := os.LookupEnv("DBName"); found {
		res.DBName = val
	}

	if val, found := os.LookupEnv("Secret"); found {
		res.Secret = val
	}
	return res
}
