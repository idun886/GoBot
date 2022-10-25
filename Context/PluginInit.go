package Context

import "github.com/gorilla/websocket"

func PluginRegister(Context string, Conn *websocket.Conn) Message {
	message := Message{Conn: Conn, Context: Context}
	return message
}
