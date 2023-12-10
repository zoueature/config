package config

type CacheConfig struct {
	Type     string `json:"type" yaml:"type"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	Debug    bool   `json:"debug" yaml:"debug"`
}
