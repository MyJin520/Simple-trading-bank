package mytools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var Encrypt *SymmetricEncryption // 全局加密器

// SymmetricEncryption AES 对称加密
type SymmetricEncryption struct {
	key string
}

func Init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *SymmetricEncryption {
	return &SymmetricEncryption{}
}

// PadPwd 填充密码长度 (PKCS7填充)
func PadPwd(srcByte []byte, blockSize int) []byte {
	PadNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(PadNum)}, PadNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// UnpadPwd 去除填充
func UnpadPwd(srcByte []byte) ([]byte, error) {
	if len(srcByte) == 0 {
		return nil, errors.New("empty data")
	}
	PadNum := int(srcByte[len(srcByte)-1])
	if PadNum > len(srcByte) || PadNum == 0 {
		return nil, errors.New("invalid padding")
	}
	for i := len(srcByte) - 1; i > len(srcByte)-PadNum-1; i-- {
		if int(srcByte[i]) != PadNum {
			return nil, errors.New("invalid padding")
		}
	}
	return srcByte[:len(srcByte)-PadNum], nil
}

// MyEncrypt AES-CBC加密
func (s *SymmetricEncryption) MyEncrypt(plaintext string) (string, error) {
	key := []byte(s.key)
	plaintextBytes := []byte(plaintext)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	// 填充明文
	plaintextBytes = PadPwd(plaintextBytes, blockSize)

	// 创建随机IV
	ciphertext := make([]byte, blockSize+len(plaintextBytes))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintextBytes)

	// 返回Base64编码的字符串
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// MyDecrypt AES-CBC解密
func (s *SymmetricEncryption) MyDecrypt(ciphertext string) (string, error) {
	key := []byte(s.key)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(ciphertextBytes) < blockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:blockSize]
	ciphertextBytes = ciphertextBytes[blockSize:]

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	// 去除填充
	plaintextBytes, err := UnpadPwd(ciphertextBytes)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

func (k *SymmetricEncryption) SetKey(key string) {
	k.key = key
}
