package mysql_gorm

import (
	"context"
	"errors"

	enUtils "github.com/IbnAnjung/movie_fest/entity/utils"
	"gorm.io/gorm"
)

type gormUnitOfWork struct {
	db *gorm.DB
}

const transactionGormUnitOfWork = "transaction-gorm-tx"

func NewGormUnitOfWork(db *gorm.DB) enUtils.UnitOfWork {
	return gormUnitOfWork{
		db: db,
	}
}

func (unit gormUnitOfWork) Begin(ctx context.Context) context.Context {
	if txDb := getTxSession(ctx); txDb != nil {
		return ctx
	}

	txDB := unit.db.Begin()
	return context.WithValue(ctx, transactionGormUnitOfWork, txDB)
}

func (unit gormUnitOfWork) Commit(ctx context.Context) error {
	txDB := getTxSession(ctx)
	if txDB == nil {
		return errors.New("failed get transcation context")
	}

	return txDB.Commit().Error
}

func (unit gormUnitOfWork) Rollback(ctx context.Context) error {
	txDB := getTxSession(ctx)
	if txDB == nil {
		return errors.New("failed get transcation context")
	}

	return txDB.Rollback().Error
}

func getTxSession(ctx context.Context) *gorm.DB {
	txDB, ok := ctx.Value(transactionGormUnitOfWork).(*gorm.DB)
	if !ok {
		return nil
	}

	return txDB
}

func getTxSessionDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	DB := db
	if txDb := getTxSession(ctx); txDb != nil {
		DB = txDb
	}

	return DB
}
