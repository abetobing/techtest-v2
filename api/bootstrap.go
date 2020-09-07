package api

import (
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
