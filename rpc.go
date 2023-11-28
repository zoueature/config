package config

type RpcType string

const (
	RpcTypeGrpc RpcType = "grpc"
	RpcTypeHttp RpcType = "http"
)

type RpcConf struct {
	Type RpcType `json:"type" yaml:"type"`
	Host string  `json:"host" yaml:"host"`
	Port string  `json:"port" yaml:"port"`
}

func (c RpcConf) CallHost() string {
	h := c.Host + ":" + c.Port
	switch c.Type {
	case RpcTypeHttp:
		return "http://" + h
	case RpcTypeGrpc:
		return h
	default:
		panic("undefined rpc type")
	}
}
