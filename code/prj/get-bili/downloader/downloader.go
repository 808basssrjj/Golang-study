package downloader

import (
	"encoding/json"
	"io"
	"net/http"
)

//InfoRequest 请求id
type InfoRequest struct {
	Bvids []string
}

//VideoInfo 接收数据
type VideoInfo struct {
	Code int `json:"code"`
	Data struct {
		Bvid  string `json:"bvid"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"data"`
}

//InfoResponse 返回数据
type InfoResponse struct {
	Infos []VideoInfo
}

// BatchDownLoadVideoInfo 下载
func BatchDownLoadVideoInfo(request InfoRequest) (InfoResponse, error) {
	var response InfoResponse

	for _, bvid := range request.Bvids {
		var videoInfo VideoInfo

		resp, err := http.Get("https://api.bilibili.com/x/web-interface/view?bvid=" + bvid)
		if err != nil {
			return InfoResponse{}, err
		}

		respbytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return InfoResponse{}, err

		}

		if err = json.Unmarshal(respbytes, &videoInfo); err != nil {
			return InfoResponse{}, err
		}

		if err := resp.Body.Close(); err != nil {
			return InfoResponse{}, err
		}

		response.Infos = append(response.Infos, videoInfo)

	}
	return response, nil
}
