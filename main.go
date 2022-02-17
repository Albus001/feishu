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
	//err1 := feishu.SendMsg("chat_id", "oc_xxxxxxxxxx", "text","Hello World")
	//if err1 != nil {
	//	fmt.Println(err1)
	//}

	msg := "{\"zh_cn\": {\"title\": \"测试\", \"content\": [[{\"tag\": \"text\", \"text\": \"你好\"}, {\"tag\": \"a\", \"text\": \"123\", \"href\": \"http://www.feishu.cn\"}]]}}"
	err2 := feishu.SendMsg("user_id", userId, "post", msg)
	if err2 != nil {
		fmt.Println(err2)
	}
	err3 := feishu.SendMsg("user_id", userId, "text", msg)
	if err3 != nil {
		fmt.Println(err3)
	}
}
