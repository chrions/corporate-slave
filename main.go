package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "https://openspeech.bytedance.com/api/v1/tts"
	token  = "T64gk8YQWlg-EGYm9u1Nj8rJkIJuF-Rp"
	appid  = "7653308108"
)

type TTSRequest struct {
	App struct {
		AppID   string `json:"appid"`
		Token   string `json:"token"`
		Cluster string `json:"cluster"`
	} `json:"app"`
	User struct {
		UID string `json:"uid"`
	} `json:"user"`
	Audio struct {
		VoiceType  string  `json:"voice_type"`
		Encoding   string  `json:"encoding"`
		SpeedRatio float64 `json:"speed_ratio"`
	} `json:"audio"`
	Request struct {
		ReqID     string `json:"reqid"`
		Text      string `json:"text"`
		Operation string `json:"operation"`
	} `json:"request"`
}

type TTSResponse struct {
	ReqID     string `json:"reqid"`
	Code      int    `json:"code"`
	Operation string `json:"operation"`
	Message   string `json:"message"`
	Sequence  int    `json:"sequence"`
	Data      string `json:"data"`
	Addition  struct {
		Duration string `json:"duration"`
	} `json:"addition"`
}

func main() {
	// 构造请求数据
	ttsReq := TTSRequest{}
	ttsReq.App.AppID = appid
	ttsReq.App.Token = token
	ttsReq.App.Cluster = "volcano_tts"
	ttsReq.User.UID = "uid123"
	ttsReq.Audio.VoiceType = "ICL_zh_female_xingganyujie_tob"
	ttsReq.Audio.Encoding = "mp3"
	ttsReq.Audio.SpeedRatio = 1.0
	ttsReq.Request.ReqID = "123456"
	ttsReq.Request.Text = "嗨！很高兴和你聊天，我会一直在这里陪你"
	ttsReq.Request.Operation = "query"

	// 转换成 JSON
	jsonData, err := json.Marshal(ttsReq)
	if err != nil {
		fmt.Println("JSON Marshal Error:", err)
		return
	}

	// 发送 HTTP 请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request Creation Error:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer;"+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Response Error:", err)
		return
	}

	// 解析 JSON 响应
	var ttsResp TTSResponse
	err = json.Unmarshal(body, &ttsResp)
	if err != nil {
		fmt.Println("JSON Unmarshal Error:", err)
		return
	}

	// 检查 API 返回状态
	if ttsResp.Code != 3000 {
		fmt.Println("API Error:", ttsResp.Message)
		return
	}

	// 解码 base64 音频数据
	audioData, err := base64.StdEncoding.DecodeString(ttsResp.Data)
	if err != nil {
		fmt.Println("Base64 Decode Error:", err)
		return
	}

	// 保存到文件
	filePath := fmt.Sprintf("/Users/fyf-mac/Desktop/voice/%s.mp3", ttsReq.Audio.VoiceType)
	err = ioutil.WriteFile(filePath, audioData, 0644)
	if err != nil {
		fmt.Println("File Write Error:", err)
		return
	}

	fmt.Println("音频文件已保存到:", filePath)
}
