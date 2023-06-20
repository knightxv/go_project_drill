package db

import (
	"fmt"
	"time"

	"github.com/knightxv/go-project-drill/pkg/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDB struct {
	//sync.RWMutex
	db *gorm.DB
}

type Writer struct{}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func initMysqlDB() {
	fmt.Println("init mysqlDB start")
	//When there is no open IM database, connect to the mysql built-in database to create openIM database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		config.Config.Mysql.DBUserName, config.Config.Mysql.DBPassword, config.Config.Mysql.DBAddress[0], "mysql")
	var db *gorm.DB
	var err1 error
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		fmt.Println("Open failed ", err.Error(), dsn)
	}
	if err != nil {
		time.Sleep(time.Duration(30) * time.Second)
		db, err1 = gorm.Open(mysql.Open(dsn), nil)
		if err1 != nil {
			fmt.Println("Open failed ", err1.Error(), dsn)
			panic(err1.Error())
		}
	}

	//Check the database and table during initialization
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", config.Config.Mysql.DBDatabaseName)
	fmt.Println("exec sql: ", sql, " begin")
	err = db.Exec(sql).Error
	if err != nil {
		fmt.Println("Exec failed ", err.Error(), sql)
		panic(err.Error())
	}
	fmt.Println("exec sql: ", sql, " end")
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		config.Config.Mysql.DBUserName,
		config.Config.Mysql.DBPassword,
		config.Config.Mysql.DBAddress[0],
		config.Config.Mysql.DBDatabaseName)

	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             time.Duration(config.Config.Mysql.SlowThreshold) * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.LogLevel(config.Config.Mysql.LogLevel),                       // Log level
			IgnoreRecordNotFoundError: true,                                                                // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                                                // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("Open failed ", err.Error(), dsn)
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Config.Mysql.DBMaxLifeTime))
	sqlDB.SetMaxOpenConns(config.Config.Mysql.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Config.Mysql.DBMaxIdleConns)

	fmt.Println("open mysql ok ", dsn)
	db.AutoMigrate(
		&UserScore{},
		&RewardEventLogs{},
		&WithdrawEventLogs{},
		&ChainUpEventLogs{},
	)
	db.Set("gorm:table_options", "CHARSET=utf8")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	if !db.Migrator().HasTable(&UserScore{}) {
		fmt.Println("CreateTable UserScore")
		db.Migrator().CreateTable(&UserScore{})
	}
	if !db.Migrator().HasTable(&RewardEventLogs{}) {
		fmt.Println("CreateTable RewardEventLogs")
		db.Migrator().CreateTable(&RewardEventLogs{})
	}
	if !db.Migrator().HasTable(&WithdrawEventLogs{}) {
		fmt.Println("CreateTable WithdrawEventLogs")
		db.Migrator().CreateTable(&WithdrawEventLogs{})
	}
	if !db.Migrator().HasTable(&ChainUpEventLogs{}) {
		fmt.Println("CreateTable ChainUpEventLogs")
		db.Migrator().CreateTable(&ChainUpEventLogs{})
	}

	DB.MysqlDB.db = db
}

func (m *mysqlDB) DefaultGormDB() *gorm.DB {
	return DB.MysqlDB.db
}
