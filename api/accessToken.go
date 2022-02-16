package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type FeiShuConf struct {
	Appid     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type AccessResponse struct {
	Code        int    `json:"code"`
	Expire      int    `json:"expire"`
	Msg         string `json:"msg"`
	AccessToken string `json:"tenant_access_token"`
}

// NewAccess 获取飞书的access_token
func (feishu *FeiShuConf) NewAccess() (accessToken string, err error) {
	var assResp AccessResponse
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	contentType := "application/json; charset=utf-8"
	body, err := json.Marshal(feishu)
	if err != nil {
		return "", err
	}
	res, err := http.Post(url, contentType, strings.NewReader(string(body)))
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &assResp)
	if err != nil {
		return "", err
	}
	if assResp.Code == 0 {
		return assResp.AccessToken, nil
	} else {
		msg := fmt.Sprintf("获取token失败，code:%d, msg:%s", assResp.Code, assResp.Msg)
		return "", errors.New(msg)
	}
}
