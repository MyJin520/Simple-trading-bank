package mytools

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// Logger 全局日志记录器，使用 logrus 库实现
var Logger *logrus.Logger

// Init 初始化日志记录器，设置日志输出文件
func init() {
	// 获取日志输出文件*
	src, err := setOutputFile()
	if err != nil {
		log.Fatalf("Failed to set log output file: %v", err)
	}
	// 若全局日志记录器还未初始化
	if Logger == nil {
		// 初始化 Logger 实例
		Logger = logrus.New()
	}
	// 设置日志输出文件
	Logger.Out = src
	// 设置日志级别为 Debug 级别，会输出所有级别的日志信息
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// setOutputFile 设置日志输出文件，若文件不存在则创建
func setOutputFile() (*os.File, error) {
	// 获取当前时间
	now := time.Now()
	// 日志文件路径
	logFilePath := ""
	// os.Getwd() --> 获取当前工作目录
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	// 检查日志目录是否存在
	_, err := os.Stat(logFilePath)
	// 若日志目录不存在
	if os.IsNotExist(err) {
		// 创建日志目录，权限为 0777
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			// 若创建失败，打印错误信息并退出程序
			log.Fatalf("err: %v", err.Error())
			return nil, err
		}
	}
	// 日志文件名，格式为 YYYY-MM-DD.log
	logFileName := now.Format("2006-01-02") + ".log"
	// 完整的日志文件路径
	fileName := logFilePath + logFileName
	// 检查日志文件是否存在
	if _, err := os.Stat(fileName); err != nil {
		// 若日志文件不存在，创建新的日志文件
		if _, err := os.Create(fileName); err != nil {
			// 若创建失败，打印错误信息并退出程序
			log.Fatalf("err: %v", err.Error())
			return nil, err
		}
	}
	// 以追加写入的模式打开日志文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		// 若打开失败，打印错误信息并退出程序
		log.Fatalf("err: %v", err.Error())
		return nil, err
	}
	return src, nil
}
