package main

import (
	"awesome-blog/cmd"
	"awesome-blog/internal/models"
)

func main() {
	cmd.Execute()
}

func init() {
	models.InitDB("mysql", "root:password@/demo?charset=utf8&parseTime=True&loc=Local")
}
