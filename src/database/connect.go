package database

import (
	// 追加でインポート
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
var err error

// DBへの接続
func SqlConnect() (database *gorm.DB) {
	DBMS := os.Getenv("DB_MS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// 作成例: go_test:password@tcp(db:3306)/go_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo

	count := 0
	Db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 3 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			Db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return Db
}
