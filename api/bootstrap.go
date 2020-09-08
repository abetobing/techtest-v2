package api

import (
	"customer/api/utils"
	"database/sql"
	"log"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

var db *sql.DB
var rds *redis.Client
var config *Config

type Config struct {
	Database  string `yml:"database"`
	RedisHost string `yml:"redis_host"`
}

func LoadConfig(cfg *Config) error {
	f, err := os.Open("config.yml")
	err = yaml.NewDecoder(f).Decode(&config)
	f.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDatabase() *sql.DB {
	db, err := sql.Open("postgres", config.Database)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetRedis() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr: config.RedisHost,
	})
	if rds.Ping().Err() != nil {
		log.Fatal(rds.Ping().Err())
	}
	return rds
}

func init() {
	err := LoadConfig(nil)
	if err != nil {
		log.Fatal("Cannot load config")
	}

	db := GetDatabase()
	salt := utils.Salt()
	username := "admin"
	password := utils.Encrypt("admin", salt)

	var existing int
	err = db.QueryRow("SELET id FROM administrators WHERE id = '1'").Scan(&existing)
	if err != nil {
		_, err := db.Exec("INSERT INTO administrators(username, password, salt) VALUES($1, $2, $3)", username, password, salt)
		log.Print(err)
		if err != nil {
			log.Fatal("Cannot initialize default admin user.", err)
		}
	}

}
