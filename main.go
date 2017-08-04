package main

import (
	proto "code.xxxxxx.com/micro/proto/scores"

	"code.xxxxxx.com/micro/scores-srv/handler"
	"github.com/micro/go-log"

	"code.xxxxxx.com/micro/common"
	"code.xxxxxx.com/micro/scores-srv/store"
	"code.xxxxxx.com/micro/scores-srv/models"
	"github.com/spf13/viper"
	"github.com/kardianos/osext"
)

func readConfig() {
	path, _ := osext.ExecutableFolder()
	viper.AddConfigPath("/etc/scores")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetDefault("addr", ":8080")
	viper.SetDefault("database_driver", "mysql")
	viper.SetDefault("database_datasource", "root@tcp(127.0.0.1:3306)/scores?charset=utf8&parseTime=True")

	viper.ReadInConfig()
}

func main() {

	readConfig()

	service := common.NewService("latest", "scores", common.NilInit)

	store, err := store.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := models.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	proto.RegisterScoreHandler(service.Server(), handler.New(store, db))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
