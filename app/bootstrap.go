package app

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/movie_fest/driver"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm"
	movieUC "github.com/IbnAnjung/movie_fest/usecase/movie"
	"github.com/IbnAnjung/movie_fest/utils"
	"github.com/google/uuid"
)

func Start(ctx context.Context) (func(), error) {

	conf, err := LoadConfig()
	if err != nil {
		return func() {}, fmt.Errorf("load config failed: %s", err.Error())
	}

	mysqlConf := conf.Mysql

	sqlConfig := driver.LoadMySqlConfig(
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Schema,
		mysqlConf.MaxLifeConnection,
		mysqlConf.MaxConnection,
		mysqlConf.MaxIddleConnection,
	)

	dbconn, mysqlCleanup, err := driver.NewMysqlConnection(ctx, sqlConfig)
	if err != nil {
		return func() {}, err
	}

	redisConf := conf.Redis
	redisConfig := driver.LoadRedisConfig(redisConf.Host, redisConf.Port, redisConf.User, redisConf.Password, redisConf.Db)
	redisConn, redisCleanup, err := driver.NewRedisConnection(ctx, redisConfig)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	_ = utils.NewRedisCaching(redisConn)

	// utils
	gormDb, err := utils.NewGormOrm(conf.App.Mode, "mysql", dbconn)
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	uof := mysql_gorm.NewGormUnitOfWork(gormDb)
	storage := utils.NewLocalStorage(conf.Http.Host, "videos")
	stringGenerator := utils.NewStringGenerator(uuid.New())
	validator, err := utils.NewValidator()
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	// repository
	movieRepository := mysql_gorm.NewMovieRepository(gormDb)
	movieGenresRepository := mysql_gorm.NewMovieGenresRepository(gormDb)
	movieHasGenresRepository := mysql_gorm.NewMovieHasGenresRepository(gormDb)

	// validator
	_, err = utils.NewValidator()
	if err != nil {
		return func() {
			mysqlCleanup()
		}, err
	}

	_ = utils.NewBycrypt()

	// usecase
	movieUsecase := movieUC.NewMovieUC(uof, storage, validator, stringGenerator, movieRepository, movieGenresRepository, movieHasGenresRepository)

	// router
	router := LoadGinRouter(*conf, movieUsecase)

	httpCleanup, err := driver.RunGinHttpServer(ctx, router, driver.LoadHttpConfig(conf.Http.Port))
	if err != nil {
		return func() {
			redisCleanup()
			mysqlCleanup()
		}, err
	}

	return func() {
		mysqlCleanup()
		redisCleanup()
		httpCleanup()
	}, nil
}
