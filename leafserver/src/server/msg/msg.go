package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

var (
	JSONProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
	Name string
}

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	// 我们也可以使用 ProtobufProcessor 注册 protobuf 消息（同时注意修改配置文件 conf/conf.go 中的 Encoding）
	JSONProcessor.Register(&Hello{})
}
