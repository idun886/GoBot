package Context

import (
	"strconv"
)

//[CQ:image,file="+getsetuapi("https://api.xhofe.top/img/api/?type=json")+",id=40000]
//[CQ:record,file=http://baidu.com/1.mp3]

type CQCode struct {
}

// 声音cq码
func Record(path string) string {
	return "[CQ:record,file=" + path + "]"
}

// 视频cq码
func Video(path string) string {
	return "[CQ:video,file=" + path + "]"
}

// @别人的cq码
func AT(qq int) string {
	qqstr := strconv.Itoa(qq)
	return "[CQ:at,qq=" + qqstr + "]"
}

// 链接分享cq码
// [CQ:share,url=http://baidu.com,title=百度]
func Share(url string, title string) string {
	return "[CQ:share,url=" + url + ",title=" + title + "]"
}

// 推荐好友或群聊 cqcode [CQ:contact,type=qq,id=10001000]
func Contact(cqtype string, qqnum int) string {
	qqnumstr := strconv.Itoa(qqnum)
	return "[CQ:contact,type=" + cqtype + ",id=" + qqnumstr + "]"
}

// 图片cqcode [CQ:image,file=http://baidu.com/1.jpg,type=show,id=40004]
func Image(path string, id int) string {
	idstr := strconv.Itoa(id)
	return "[CQ:image,file=" + path + ",id=" + idstr + "]"
}

// 回复 cqcode [CQ:reply,text=Hello World,qq=10086,time=3376656000,seq=5123]
func Reply(text string, qq int, time int, messageid int, seq int) string {
	seqstr := strconv.Itoa(seq)
	messageidstr := strconv.Itoa(messageid)
	qqstr := strconv.Itoa(qq)
	timestr := strconv.Itoa(time)
	return "[CQ:reply" + ",seq=" + seqstr + ",text=" + text + ",id=" + messageidstr + ",qq=" + qqstr + ",time=" + timestr + "]"
}

// [CQ:reply,id=123456]
func Replytest(messageid int) string {
	messageidstr := strconv.Itoa(messageid)
	return "[CQ:reply,id=" + messageidstr + "]"
}
