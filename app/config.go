package app

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Mode string
}

type JwtConfig struct {
	SecretKey      string
	ExpireDuration int64
}

type Config struct {
	App   AppConfig
	Http  HttpConfig
	Mysql MysqlConfig
	Redis RedisConfig
	Jwt   JwtConfig
}

type HttpConfig struct {
	Host string
	Port int
}

type MysqlConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	Schema             string
	MaxLifeConnection  int64
	MaxConnection      int
	MaxIddleConnection int
}

type RedisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	httpConfig, err := loadHttpConfig()
	if err != nil {
		return nil, err
	}

	mysqlConfig, err := loadMySqlConfig()
	if err != nil {
		return nil, err
	}

	redisConfig, err := loadRedisConfig()
	if err != nil {
		return nil, err
	}

	jwtConfig, err := loadJwtConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Mode: os.Getenv("APP_MODE"),
		},
		Http:  httpConfig,
		Mysql: mysqlConfig,
		Redis: redisConfig,
		Jwt:   jwtConfig,
	}, nil
}

func loadHttpConfig() (HttpConfig, error) {
	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return HttpConfig{}, errors.New("invalid HTTP_PORT value")
	}

	return HttpConfig{
		Host: os.Getenv("HTTP_HOST"),
		Port: httpPort,
	}, nil
}

func loadMySqlConfig() (MysqlConfig, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return MysqlConfig{}, errors.New("invalid DB_PORT value")
	}

	maxLifeConnection, err := strconv.ParseInt(os.Getenv("DB_MAX_LIFE_CONNECTION"), 10, 64)
	if err != nil {
		return MysqlConfig{}, errors.New("invalid DB_MAX_LIFE_CONNECTION value")
	}

	if maxLifeConnection == 0 {
		maxLifeConnection = 180
	}

	maxConnection, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION"))
	if err != nil {
		return MysqlConfig{}, errors.New("invalid DB_MAX_CONNECTION value")
	}

	if maxConnection == 0 {
		maxConnection = 10
	}

	maxIddleConnection, err := strconv.Atoi(os.Getenv("DB_MAX_IDDLE_CONNECTION"))
	if err != nil {
		return MysqlConfig{}, errors.New("invalid DB_MAX_IDDLE_CONNECTION value")
	}

	if maxIddleConnection == 0 {
		maxIddleConnection = 10
	}

	return MysqlConfig{
		Host:               os.Getenv("DB_HOST"),
		Port:               port,
		User:               os.Getenv("DB_USER"),
		Password:           os.Getenv("DB_PASSWORD"),
		Schema:             os.Getenv("DB_SCHEMA"),
		MaxLifeConnection:  maxLifeConnection,
		MaxConnection:      maxConnection,
		MaxIddleConnection: maxIddleConnection,
	}, nil
}

func loadRedisConfig() (RedisConfig, error) {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return RedisConfig{}, errors.New("invalid REDIS_PORT value")
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return RedisConfig{}, errors.New("invalid REDIS_DB value")
	}

	return RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     port,
		User:     os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Db:       db,
	}, nil
}

func loadJwtConfig() (JwtConfig, error) {
	expireDuration, err := strconv.ParseInt(os.Getenv("JWT_EXIPIRE_DURATION"), 10, 64)
	if err != nil {
		return JwtConfig{}, errors.New("invalid HTTP_PORT value")
	}

	return JwtConfig{
		SecretKey:      os.Getenv("JWT_KEY"),
		ExpireDuration: expireDuration,
	}, nil
}
