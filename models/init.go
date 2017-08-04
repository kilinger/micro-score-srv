package models

import (
    "log"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/spf13/viper"
)

var (
    driver = "mysql"
    path = "root@tcp(127.0.0.1:3306)/scores?charset=utf8&parseTime=True"
)

func GetDB() (*gorm.DB, error) {

	db, err := gorm.Open(viper.GetString("database_driver"), viper.GetString("database_datasource"))

	if err != nil {
        fmt.Println(err)
        log.Fatal(err)
		return nil, err
	}

	db.LogMode(true)
	Migrate(db)

	return db, nil
}