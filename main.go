package main

import (
	"CloudStorage/apps"
	"CloudStorage/conf"
	csLog "CloudStorage/log"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("CloudStorage initialization Panic: ", err)
		}
	}()

	if err := conf.ICloudConfInit(); err != nil {
		log.Fatal("CloudStorage Configuration Initialize Error: ", err)
	}

	csLog.InitLogger()
}

func main() {
	r := gin.Default()

	apps.RouterSetUp(r)

	webServerAddr := conf.CSConf.Ip + ":" + strconv.Itoa(conf.CSConf.Port)
	if err := r.Run(webServerAddr); err != nil {
		log.Println("gin server start error:", err)
	}
}