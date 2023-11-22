package config

type AppConfig struct {
	Listen string `json:"listen" yaml:"listen"`
	Debug  bool   `json:"debug" yaml:"debug"`
}
