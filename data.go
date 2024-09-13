package main

import (
	"github.com/cavaliergopher/grab/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

// 单例结构体

type DataManager struct {
	data  EPG
	mutex sync.Mutex
}

// 全局单例变量

var instance *DataManager
var once sync.Once

// 获取单例实例

func GetDataManagerInstance() *DataManager {
	once.Do(func() {
		instance = &DataManager{
			data: makeData(),
		}
		// 启动一个协程定时更新数据
		go instance.updateDataPeriodically()
	})
	return instance
}

// 定时更新数据的函数

func (dm *DataManager) updateDataPeriodically() {
	ticker := time.NewTicker(2 * time.Hour) // 每5秒更新一次数据
	for {
		select {
		case <-ticker.C:
			dm.mutex.Lock()
			dm.data = makeData()
			dm.mutex.Unlock()
		}
	}
}

// 获取数据

func (dm *DataManager) GetData() EPG {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()
	return dm.data
}

func makeData() EPG {
	log.Debugf("开始更新XML文件,时间 - %s", time.Now().String())
	var file = "epg.xml"
	//删除老的文件
	_ = os.Remove(file)
	resp, err := grab.Get(file, viper.GetString("epg_url"))
	if err != nil {
		log.Error("下载XML文件失败", err)
		return EPG{}
	}
	return parseXML(resp.Filename)
}
