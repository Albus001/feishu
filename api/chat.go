package api

import (
	"encoding/json"
	"errors"
	"feishu/common"
	"fmt"
)

type ChatResponse struct {
	Code int      `json:"code"`
	Data ChatData `json:"data"`
	Msg  string   `json:"msg"`
}

type ChatData struct {
	HasMore   bool       `json:"has_more"`
	Items     []ChatItem `json:"items"`
	PageToken string     `json:"page_token"`
}

type ChatItem struct {
	ChatId string `json:"chat_id"`
	Name   string `json:"name"`
}

// GetChats 获取机器人所在的所有群组
func (feishu *FeiShu) GetChats() (chatItems []ChatItem, err error) {
	var chatResp ChatResponse
	//var resp = make(map[string]interface{})
	url := "https://open.feishu.cn/open-apis/im/v1/chats?user_id_type=user_id"
	headers := feishu.GetHeaders()
	bytes, err := common.SendGetRequest(url, headers)
	if err != nil {
		return chatItems, err
	}
	//fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &chatResp)
	if err != nil {
		return chatItems, err
	}
	if chatResp.Code == 0 {
		if len(chatResp.Data.Items) > 0 {
			chatItems = chatResp.Data.Items
			return chatItems, err
		} else {
			msg := "没有发现机器人所在群组，请先添加机器人到群组"
			return chatItems, errors.New(msg)
		}
	} else {
		msg := fmt.Sprintf("获取机器人所在群组失败，code:%d, msg:%s", chatResp.Code, chatResp.Msg)
		return chatItems, errors.New(msg)
	}
}

func (feishu *FeiShu) GetChatByName(chatName string) (chatId string, err error) {
	chatItems, err := feishu.GetChats()
	if err != nil {
		return chatId, err
	}
	var chatNameTemp string
	length := len(chatItems)
	for index, chatItem := range chatItems {
		chatNameTemp += fmt.Sprintf("'%s'", chatItem.Name)
		if index < length-1 {
			chatNameTemp += ","
		}
		if chatItem.Name == chatName {
			chatId = chatItem.ChatId
			return chatId, nil
		}
	}
	msg := fmt.Sprintf("没有找到指定的群组，可选的群组名称为:%s", chatNameTemp)
	return chatId, errors.New(msg)

	//if userInfoResp.Code == 0 {
	//	if len(userInfoResp.Data.UserList) > 0 {
	//		userId = userInfoResp.Data.UserList[0].UserId
	//		if userId != "" {
	//			return userId, nil
	//		} else {
	//			msg := "获取用户信息失败，输入的手机号不存在"
	//			return "", errors.New(msg)
	//		}
	//	} else {
	//		msg := "获取用户信息失败，输入的手机号不存在"
	//		return "", errors.New(msg)
	//	}
	//} else {
	//	msg := fmt.Sprintf("获取用户信息失败，code:%d, msg:%s", userInfoResp.Code, userInfoResp.Msg)
	//	return "", errors.New(msg)
	//}
}
