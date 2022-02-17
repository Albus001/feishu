package main

import (
	"fmt"
	"github.com/Albus001/feishu/api"
)

func main() {
	feishuConf := api.FeiShuConf{
		"appid",
		"app_secret",
	}
	access, err := feishuConf.NewAccess()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(access)
	var feishu = api.FeiShu{access}
	chatId, err := feishu.GetChatByName("test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(chatId)
	mobile := "186xxxxxxxxx"
	userId, err := feishu.GetUserInfo("user_id", mobile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
	//err1 := feishu.SendMessage("text", userId, "Hello World")
	err1 := feishu.SendMsg("chat_id", "oc_xxxxxxxxxx", "Hello World")
	if err1 != nil {
		fmt.Println(err1)
	}
	err2 := feishu.SendMsg("user_id", userId, "Hello World")
	if err2 != nil {
		fmt.Println(err2)
	}
}
