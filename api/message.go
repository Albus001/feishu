package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Albus001/feishu/common"
)

type MessageBody struct {
	MsgType   string `json:"msg_type"`
	ReceiveId string `json:"receive_id"`
	Content   string `json:"content"`
}

type MessageBody2 struct {
	MsgType string  `json:"msg_type"`
	UserId  string  `json:"user_id"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SendMessageToUser 这个接口只能用来给用户发消息(通过user_id)
func (feishu *FeiShu) SendMessageToUser(messageType string, userId string, msg string) (err error) {
	url := "https://open.feishu.cn/open-apis/message/v4/send/"
	headers := feishu.GetHeaders()
	var messageBody MessageBody2
	var messageResp MessageResponse
	messageBody.MsgType = messageType
	messageBody.UserId = userId
	messageBody.Content = Content{msg}
	body, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}
	bytes, err := common.SendPostRequest(url, headers, body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &messageResp)
	if err != nil {
		return err
	}
	if messageResp.Code == 0 {
		return nil
	} else {
		msg = fmt.Sprintf("消息发送失败，code:%d, msg:%s", messageResp.Code, messageResp.Msg)
		return errors.New(msg)
	}
}

// SendMsg 根据receiveIdType 给用户(user_id或open_id)或群组(chat_id)发送消息
func (feishu *FeiShu) SendMsg(receiveIdType string, receiveId string, msg string) (err error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=%s", receiveIdType)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", feishu.AccessToken)
	headers["Content-Type"] = "application/json; charset=utf-8"
	var messageBody MessageBody
	var messageResp MessageResponse
	messageBody.MsgType = "text"
	messageBody.ReceiveId = receiveId
	content, _ := json.Marshal(Content{msg})
	messageBody.Content = string(content)
	body, err := json.Marshal(messageBody)
	bytes, err := common.SendPostRequest(url, headers, body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &messageResp)
	if err != nil {
		return err
	}
	if messageResp.Code == 0 {
		return nil
	} else {
		msg = fmt.Sprintf("消息发送失败，code:%d, msg:%s", messageResp.Code, messageResp.Msg)
		return errors.New(msg)
	}
}
