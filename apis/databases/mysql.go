package databases

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"smiling-blog/config"
	"smiling-blog/tools/slog"
)

type Db struct {
	db *gorm.DB
}

func (slf *Db) DB() *gorm.DB {
	if slf.db == nil {
		slf.StartDb()
	}
	return slf.db
}

func (slf *Db) StartDb() error {
	if slf.db != nil {
		return errors.New("Db already open")
	}
	driver := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.Configs.DataBase.User,
		config.Configs.DataBase.Pwd,
		config.Configs.DataBase.Host,
		config.Configs.DataBase.Port,
		config.Configs.DataBase.DbName)
	var err error
	slf.db, err = gorm.Open("mysql", driver)

	if err != nil {
		slog.Infof("conn Db  error %s", err)
		return err
	}
	slf.db.DB().SetMaxIdleConns(config.Configs.DataBase.MaxIdleCons)
	slf.db.DB().SetMaxOpenConns(config.Configs.DataBase.MaxOpenCons)
	slf.db.DB().SetConnMaxLifetime(config.Configs.DataBase.MaxLifeTime)

	slf.db.SingularTable(true)
	slf.db.LogMode(true)
	slf.db.SetLogger(&slog.Slog)
	return nil
}

func (slf *Db) StopDb() error {
	if slf.db != nil {
		err := slf.db.Close()
		if err != nil {
			slf.db = nil
		}
		return err
	}
	return errors.New("Db is nil")
}

func (slf *Db) Tx(f func(db *gorm.DB) error) error {
	tx := slf.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	if e := f(tx); e != nil {
		tx.Rollback()
		return e
	}
	if e1 := tx.Commit().Error; e1 != nil {
		return e1
	}
	return nil
}

func (slf *Db) ExecuteSql(f func(db *gorm.DB) (interface{}, error)) (interface{}, error) {
	result, ok := f(slf.db)
	if ok != nil {
		return result, ok
	}
	return result, nil
}
