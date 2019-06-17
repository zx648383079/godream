package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHash 哈希密码，统一php
func PasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// PasswordVerify 哈希密码验证
func PasswordVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Md5Str MD5 加密字符串
func Md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
