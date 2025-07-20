package service

import (
	"go-store/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

// UploadAvatarToLocalStatic 上传头像
func UploadAvatarToLocalStatic(file multipart.File, userId int, userName string) (string, error) {
	bid := strconv.Itoa(userId)
	basePath := "." + conf.AvatarPath + "user" + bid

	// 创建目录（如果不存在）
	if !isDirExist(basePath) {
		CreateDir(basePath)
	}

	// 专属与用户的存储路径--构建文件路径（防止重名，可加时间戳或随机字符串）
	t := strconv.FormatInt(time.Now().Unix(), 10)
	avatarPath := basePath + "/" + userName + t + ".jpg"

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}

	return "user" + bid + "/" + userName + t + ".jpg", nil
}

func isDirExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // 其他错误（如权限问题）
	}
	return info.IsDir() // 存在且是文件夹
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}

/* ------------商品操作-------------------- */

func UploadProductLocalStatic(file multipart.File, userId int, productName string) (string, error) {
	bid := strconv.Itoa(userId)
	basePath := "." + conf.ProductPath + "boss" + bid

	// 创建目录（如果不存在）
	if !isDirExist(basePath) {
		CreateDir(basePath)
	}

	// 专属与用户的存储路径--构建文件路径（防止重名，可加时间戳或随机字符串）
	t := strconv.FormatInt(time.Now().Unix(), 10)
	productPath := basePath + "/" + productName + t + ".jpg"

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}

	return "boss" + bid + "/" + productName + t + ".jpg", nil
}
