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
