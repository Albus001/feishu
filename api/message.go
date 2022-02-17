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

type TextContent struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// SendMsg 根据receiveIdType 给用户(user_id或open_id)或群组(chat_id)发送消息
func (feishu *FeiShu) SendMsg(receiveIdType string, receiveId string, msgType, msg string) (err error) {
	url := fmt.Sprintf("https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=%s", receiveIdType)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", feishu.AccessToken)
	headers["Content-Type"] = "application/json; charset=utf-8"
	var messageBody MessageBody
	var messageResp MessageResponse
	messageBody.MsgType = msgType
	messageBody.ReceiveId = receiveId
	if msgType == "text" {
		content, _ := json.Marshal(TextContent{msg})
		messageBody.Content = string(content)
	} else if msgType == "post" {
		messageBody.Content = msg
	} else {
		messageBody.Content = msg
	}

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
