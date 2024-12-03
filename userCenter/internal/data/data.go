package data

import (
	"fmt"
	"moonChat/userCenter/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	rdb *redis.Client
	psq *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	//创建redis实例
	rdb := NewRedisCli(c)
	psq := NewPsqlCli(c)
	d := &Data{
		rdb: rdb,
		psq: psq,
	}

	//自动更新表结构或创建表
	return d, cleanup, nil
}

func NewRedisCli(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	return rdb
}

func NewPsqlCli(c *conf.Data) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", c.Psql.User, c.Psql.Password, c.Psql.Host, c.Psql.Port, c.Psql.Dbname)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("psql fail to connect ---", err)
		return nil
	}
	return conn
}
