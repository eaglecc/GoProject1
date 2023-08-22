package main

import (
	downloader "awesomeProject/downloader"
	myfmt "awesomeProject/fmt"
)

func main() {
	request := downloader.InfoRequest{Bvids: []string{"BV1Jo4y1y7yZ", "BV1184y1o7yP", "BV1Yx4y1X7Kg"}}
	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}
	for _, info := range response.Infos {
		myfmt.Logger.Printf("title:%s \ndesc:%s\n", info.Data.Title, info.Data.Desc)
	}
}
