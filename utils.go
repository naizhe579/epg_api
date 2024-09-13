package main

import (
	"encoding/xml"
	"github.com/gofiber/fiber/v3/log"
	"os"
	"time"
)

func parseXML(filePath string) EPG {
	xmlBody, err := os.ReadFile(filePath)
	if err != nil {
		log.Error("读取XML文件失败", err.Error())
		return EPG{}
	}
	var xmlTv XMLTv
	err = xml.Unmarshal(xmlBody, &xmlTv)
	if err != nil {
		log.Error("解析XML文件失败", err.Error())
		return EPG{}
	}
	var epg = EPG{
		InfoName:    xmlTv.InfoName,
		InfoUrl:     xmlTv.InfoUrl,
		ChannelsMap: make(map[string]*Channel),
	}
	//遍历channel
	for _, xmlChannel := range xmlTv.Channels {
		item := &Channel{
			ID:         xmlChannel.ID,
			Programmes: make([]*Programme, 0),
		}
		epg.ChannelsMap[item.ID] = item
	}
	//遍历programme
	for _, xmlProgramme := range xmlTv.Programmes {
		item := &Programme{
			Channel: xmlProgramme.Channel,
			Title:   xmlProgramme.Title,
			Start:   parseTime(xmlProgramme.Start),
			Stop:    parseTime(xmlProgramme.Stop),
		}
		//找到对应的channel
		channel := epg.ChannelsMap[item.Channel]
		channel.Programmes = append(channel.Programmes, item)
	}
	return epg
}

func parseTime(data string) time.Time {
	timeLayout := "20060102150405 -0700"
	t, err := time.Parse(timeLayout, data)
	if err != nil {
		return time.Now()
	}
	return t
}

func formatTime(time time.Time) string {
	return time.Format("15:04")
}
