package config

type AppConfig struct {
	Listen     string `json:"listen" yaml:"listen"`
	GrpcListen string `json:"grpcListen" yaml:"grpcListen"`
	Debug      bool   `json:"debug" yaml:"debug"`
	Env        string `json:"env" yaml:"env"`
}
