package main

import (
	"github.com/Terence1105/LineBot/api"
	"github.com/Terence1105/LineBot/db"
	_ "github.com/Terence1105/LineBot/viper"

	"github.com/spf13/viper"
)

func main() {
	memePictures := api.New_MemesPicture()

	lineBot := api.NewLineBot()

	messageLogDB := db.NewDB()

	server := api.NewServer(memePictures, lineBot, messageLogDB)

	server.Start(":" + viper.GetString("port"))
}
