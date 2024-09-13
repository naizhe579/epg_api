package main

import (
	"fmt"
	"time"
)

//实体类

type EPG struct {
	InfoName    string              `json:"info_name"`
	InfoUrl     string              `json:"info_url"`
	ChannelsMap map[string]*Channel `json:"channels_map"`
}

type Channel struct {
	ID         string       `json:"id"`
	Programmes []*Programme `json:"programmes"`
}

type Programme struct {
	Channel string    `json:"channel"`
	Title   string    `json:"title"`
	Start   time.Time `json:"start"`
	Stop    time.Time `json:"stop"`
}

func (epg EPG) GetCurrentProgramme(channelID string, currentTime time.Time) string {
	// 从 ChannelsMap 中找到对应的频道
	channel, exists := epg.ChannelsMap[channelID]
	if !exists {
		return "暂无节目信息"
	}

	var currentProgrammeIndex = -1
	// 遍历该频道的节目列表，查找当前时间正在播放的节目
	for index, programme := range channel.Programmes {
		if currentTime.After(programme.Start) && currentTime.Before(programme.Stop) {
			currentProgrammeIndex = index
			break
		}
	}

	// 如果没有找到当前时间的节目
	if currentProgrammeIndex == -1 {
		return "暂无播放节目"
	}

	// 获取当前节目
	p1 := channel.Programmes[currentProgrammeIndex]
	s1 := fmt.Sprintf("%s - %s %s", formatTime(p1.Start), formatTime(p1.Stop), p1.Title)

	// 检查是否有下一个节目
	var s2 string
	if currentProgrammeIndex+1 < len(channel.Programmes) {
		p2 := channel.Programmes[currentProgrammeIndex+1]
		s2 = fmt.Sprintf("%s - %s %s", formatTime(p2.Start), formatTime(p2.Stop), p2.Title)
	} else {
		s2 = "暂无后续节目"
	}

	return s1 + "\n" + s2
}
