package main

import (
	"get-bili/downloader"
	myfmt "get-bili/fmt"
)

func main() {
	// fmt.Println("ss")
	// myfmt.Logger.Println("hello")

	requset := downloader.InfoRequest{Bvids: []string{"BV1Ff4y187q9", "BV1Jo4y1y7yZ"}}
	reqonse, err := downloader.BatchDownLoadVideoInfo(requset)
	if err != nil {
		panic(err)
	}
	for _, info := range reqonse.Infos {
		myfmt.Logger.Printf("title:%s\n desc:%s\n", info.Data.Title, info.Data.Desc)
	}
}
