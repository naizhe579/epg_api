package main

// XML 解析用的结构体

type XMLTv struct {
	InfoName   string         `xml:"info-name,attr"`
	InfoUrl    string         `xml:"info-url,attr"`
	Channels   []XMLChannel   `xml:"channel"`
	Programmes []XMLProgramme `xml:"programme"`
}

type XMLChannel struct {
	ID          string `xml:"id,attr"`
	DisplayName string `xml:"display-name"`
}

type XMLProgramme struct {
	Channel string `xml:"channel,attr"`
	Start   string `xml:"start,attr"`
	Stop    string `xml:"stop,attr"`
	Title   string `xml:"title"`
}
