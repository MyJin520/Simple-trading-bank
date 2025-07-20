package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDB     string
	RedisAddr   string
	RedisPwd    string
	RedisDBName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：" + err.Error())
	}
	MyLoadRedis(file)
	if err := Redis(); err != nil {
		fmt.Printf("Redis 初始化失败: %v\n", err)
		panic(err)
	}
}

func Redis() error {
	// 将 RedisDB 字符串转换为 uint64 类型
	db, err := strconv.ParseUint(RedisDB, 10, 64)
	if err != nil {
		return fmt.Errorf("解析 Redis 数据库编号失败: %w", err)
	}

	// 创建 Redis 客户端实例
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPwd, // 根据配置动态设置密码
		DB:       int(db),
	})

	_, err = client.Ping().Result()
	if err != nil {
		return fmt.Errorf("redis 连接测试失败: %w", err)
	}

	// 将初始化好的客户端赋值给全局变量
	RedisClient = client
	return nil
}

func MyLoadRedis(file *ini.File) {
	RedisDB = file.Section("redis").Key("RedisDB").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPwd = file.Section("redis").Key("RedisPw").String()
	RedisDBName = file.Section("redis").Key("RedisDbName").String()
}
