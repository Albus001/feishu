package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Albus001/feishu/common"
)

type FeiShu struct {
	AccessToken string `json:"access_token"`
}

type UserInfoBody struct {
	Mobiles []string `json:"mobiles"`
}

type UserInfoResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data UserListTemp `json:"data"`
}

type UserListTemp struct {
	UserList []UserTemp `json:"user_list"`
}

type UserTemp struct {
	Mobile string `json:"mobile"`
	UserId string `json:"user_id"`
}

func (feishu *FeiShu) GetHeaders() (headers map[string]string) {
	headers = make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", feishu.AccessToken)
	headers["Content-Type"] = "application/json; charset=utf-8"
	return headers
}

// GetUserInfo 根据手机号获取user_id
func (feishu *FeiShu) GetUserInfo(userIdType string, mobile string) (userId string, err error) {
	var userInfoResp UserInfoResponse
	var userInfoBody UserInfoBody
	userInfoBody.Mobiles = append(userInfoBody.Mobiles, mobile)
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id?user_id_type=%s", userIdType)
	headers := feishu.GetHeaders()
	body, err := json.Marshal(userInfoBody)
	if err != nil {
		return "", err
	}
	bytes, err := common.SendPostRequest(url, headers, body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &userInfoResp)
	if err != nil {
		return "", err
	}
	if userInfoResp.Code == 0 {
		if len(userInfoResp.Data.UserList) > 0 {
			userId = userInfoResp.Data.UserList[0].UserId
			if userId != "" {
				return userId, nil
			} else {
				msg := "获取用户信息失败，输入的手机号不存在"
				return "", errors.New(msg)
			}
		} else {
			msg := "获取用户信息失败，输入的手机号不存在"
			return "", errors.New(msg)
		}
	} else {
		msg := fmt.Sprintf("获取用户信息失败，code:%d, msg:%s", userInfoResp.Code, userInfoResp.Msg)
		return "", errors.New(msg)
	}
}
