package common

//客户端，服务器端数据交互模板
type EgoResult struct {
	Status int         //状态200表示成功
	Data   interface{} //返回的数据
	Msg    string      //返回的消息
}
