package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// Database 对数据库进行初始化和管理
type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}
func GetSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func openDB(username, password, addr, name string) *gorm.DB {

	mysqlConfig := fmt.Sprintf("%"+
		"s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	db, err := gorm.Open("mysql", mysqlConfig)
	if err != nil {
		log.Errorf(err, "数据库连接失败.Database Name: %s", name)
	}
	db.LogMode(viper.GetBool("run_mode"))
	db.DB().SetMaxIdleConns(0)
	db.SingularTable(true)
	db.LogMode(true)
	return db
}
