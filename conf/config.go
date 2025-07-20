package conf

import (
	"fmt"
	"go-store/dao"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DB         string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string // u大小写错了这里
	DBPassWord string

	RedisDB     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtPass    string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	// 本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMySQL(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	// 读取mysql--主
	ReadDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassWord, DBHost, DBPort, DBName)
	// 写入mysql--从
	WriteDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassWord, DBHost, DBPort, DBName)
	dao.Database(ReadDSN, WriteDSN)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
func LoadMySQL(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBName = file.Section("mysql").Key("DBName").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPassWord = file.Section("mysql").Key("DBPassWord").String()
}

func LoadRedis(file *ini.File) {
	RedisDB = file.Section("redis").Key("RedisDB").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtPass = file.Section("email").Key("SmtPass").String()
}
func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
