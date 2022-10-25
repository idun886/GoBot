package Context

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type Message struct {
	Conn           *websocket.Conn //websocket链接
	Context        string          //待解析的json报文
	Message        string          //解析后的message消息 不分群组和私聊
	MessageID      int             //消息id
	PrivateMessage string          //解析后的 私聊消息
	GroupMessage   string          //解析后的 群组消息
	UserID         int             //解析后的发送人的qq
	GroupID        int             //解析后的该消息的发送群号
}

// 获取消息的内容字符串 只解析出来 最原始message消息  如果里面有cq码则会把cq码提取出来  并且把消息字符串赋值给 结构体中的变量
func (m *Message) ExplainMessage() {
	message := gjson.Get(m.Context, "message").String()
	message_type := gjson.Get(m.Context, "message_type").String()
	user_id := gjson.Get(m.Context, "user_id").Int()
	group_id := gjson.Get(m.Context, "group_id").Int()
	message_id := gjson.Get(m.Context, "message_id").Int()
	//如果消息类型是私聊
	if message_type == "private" {
		m.UserID = int(user_id)
		m.Message = message
		m.MessageID = int(message_id)
		m.PrivateMessage = message
		//fmt.Printf("收到一条来自%d此人的私聊消息：%s\n", m.UserID, m.PrivateMessage)
		//	如果是群聊
	} else if message_type == "group" {
		m.UserID = int(user_id)
		m.GroupID = int(group_id)
		m.GroupMessage = message
		m.Message = message
		m.MessageID = int(message_id)
		//fmt.Printf("收到一条来自%d此群的消息：%s，发送人%d\n", m.GroupID, m.GroupMessage, m.UserID)
	}
	m.Message = message
}

// 获取消息方法
func (m *Message) On_Message() string {
	return m.Message
}

// 获取私聊消息方法
func (m *Message) Get_Private_Message() string {
	return m.PrivateMessage
}

// 获取群组消息方法
func (m *Message) Get_Group_Message() string {
	return m.GroupMessage
}

// 获取群号方法
func (m *Message) Get_GroupId() int {
	return m.GroupID
}

// 获取发送者qq号方法
func (m *Message) Get_UserId() int {
	return m.UserID
}

// 获取消息类型的方法
func (m *Message) Get_Message_Type() string {
	message_type := gjson.Get(m.Context, "message_type").String()
	if message_type == "private" {
		return "private"
	} else if message_type == "group" {
		return "group"
	} else if message_type == "guild" {
		return "guild"
	}
	return "nil"
}

// 发送私聊消息
func (m *Message) Send_Private_Message(content string, userid int) {
	var private_message_struct string = `{
    "action": "send_private_msg",
    "params": {
        "user_id": "` + strconv.Itoa(userid) + `",
        "message": "` + content +
		`"
    },
    "echo": "发送一条私聊消息"
	}`
	m.Conn.WriteMessage(websocket.TextMessage, []byte(private_message_struct))

}

// 发送群聊消息
func (m *Message) Send_Group_Message(content string, groupid int) {
	var group_message_struct string = `{
    "action": "send_group_msg",
    "params": {
        "group_id": "` + strconv.Itoa(groupid) + `",
        "message": "` + content + `"
    },
    "echo": "发送一条群消息"
	}`
	m.Conn.WriteMessage(websocket.TextMessage, []byte(group_message_struct))
}

// 根据特定消息进行回复   可以定义规则
func (m *Message) On_Commend(rule string, content string) {
	//如果是私聊
	if rule == m.Message && m.Get_Message_Type() == "private" {
		userid := strconv.Itoa(m.Get_UserId())
		var send_message_struct string = `{
		"action": "send_private_msg",
    	"params": {
        "user_id": "` + userid + `",
        "message": "` + content + `"
    	},
    	"echo": "发送一条私聊消息"
		}`
		m.Conn.WriteMessage(websocket.TextMessage, []byte(send_message_struct))
		fmt.Println("Reply私聊方法执行")
	}

	if rule == m.Message && m.Get_Message_Type() == "group" {
		groupid := strconv.Itoa(m.Get_GroupId())
		var send_message_struct string = `{
    	"action": "send_group_msg",
    	"params": {
        	"group_id": "` + groupid + `",
        	"message": "` + content + `"
		},
    	"echo": "发送一条群消息"
		}`
		m.Conn.WriteMessage(websocket.TextMessage, []byte(send_message_struct))
		fmt.Println("Reply群聊方法执行")
	}

}

// 关键字匹配  如果 接收到的消息里 有定义的关键字 就触发函数
func (m *Message) On_Keyword(keyword string, context string) {
	//判断 接收到的消息里是否包含keyword
	findbool := strings.Contains(m.Message, keyword)
	if findbool && m.Get_Message_Type() == "private" {
		userid := strconv.Itoa(m.Get_UserId())
		var send_message_struct string = `{
				"action": "send_private_msg",
    			"params": {
        		"user_id": "` + userid + `",
        		"message": "` + context +
			`"
    			},
    			"echo": "发送一条私聊消息"
			}`
		m.Conn.WriteMessage(websocket.TextMessage, []byte(send_message_struct))
	}
	if findbool && m.Get_Message_Type() == "group" {
		groupid := strconv.Itoa(m.Get_GroupId())
		var send_message_struct string = `{
			"action": "send_group_msg",
			"params": {
        		"group_id": "` + groupid + `",
        		"message": "` + context + `"
			},
    			"echo": "发送一条群消息"
			}`
		m.Conn.WriteMessage(websocket.TextMessage, []byte(send_message_struct))
	}

}
