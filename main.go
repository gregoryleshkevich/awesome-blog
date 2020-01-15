package main

import (
	"awesome-blog/cmd"
	"awesome-blog/internal/models"
	"awesome-blog/internal/services"
)

func main() {
	cmd.Execute()
}

func init() {
	services.InitConfig()
	models.InitDB(services.Config.DBName, services.Config.DataSourceName)
}
