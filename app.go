package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//默认配置
	viper.SetDefault("addr", ":9000")
	viper.SetDefault("epg_url", "https://epg.112114.xyz/pp.xml")
	//读取配置文件
	viper.SetConfigName("epg_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	addr := viper.GetString("addr")
	epgUrl := viper.GetString("epg_url")
	log.Debugf("addr: \"%s\" - epg_url: \"%s\"", addr, epgUrl)
	//启动数据库初始任务
	log.Debugf("epg: %v", GetDataManagerInstance().GetData())
	//启动WEB服务
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/getInfo", func(c fiber.Ctx) error {
		var channel = c.Query("channel")
		//这里可以做个容错处理，名称别名自动转成标准的EPG名称
		var data = GetDataManagerInstance().GetData().GetCurrentProgramme(channel, time.Now())
		return c.SendString(data)
	})
	log.Fatal(app.Listen(addr))
}
