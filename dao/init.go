package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		// 写入debug日志
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string类型字段默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度 ，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引，就要把索引先删除再重建，mysql5.7 不支持
		DontSupportRenameColumn:   true, //用change重命名列 ，mysql8之前的数据库不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表命名策略，表名不加s
		},
	})

	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 最大连接数
	sqlDB.SetMaxOpenConns(100) // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      // 写操作
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作
		Policy:   dbresolver.RandomPolicy{},
	}))

	// 数据库迁移
	Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
