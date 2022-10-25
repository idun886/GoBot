package Context

import (
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

type Meta struct {
	Context  string
	Conn     *websocket.Conn
	MetaType string
	PostType string
	SelfID   int
}

func (m *Meta) Login() {

	posttype := gjson.Get(m.Context, "post_type").String()
	metatype := gjson.Get(m.Context, "meta_event_type").String()
	selfid := gjson.Get(m.Context, "self_id").Int()
	m.SelfID = int(selfid)
	m.MetaType = metatype
	m.PostType = posttype

}


