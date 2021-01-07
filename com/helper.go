package com

import (
	"github.com/astaxie/beego/config"
	"path/filepath"
	"os"
	"log"
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func AppPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func AppConfig() (config.Configer, error) {
	return config.NewConfig("ini", "conf/app.conf")
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandStringWithLetters(n int, letters string) string {
	myLetterRunes := []rune(letters)
	b := make([]rune, n)
	for i := range b {
		b[i] = myLetterRunes[rand.Intn(len(myLetterRunes))]
	}
	return string(b)
}

func MD5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}
