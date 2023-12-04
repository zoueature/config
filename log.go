package config

const (
	StdoutDriver = "stdout"
	FileDriver   = "file"
)

type LogConfig struct {
	Driver            string `json:"driver" yaml:"driver"`
	Path              string `json:"path" yaml:"path"`
	MinLevel          string `json:"minLevel"`
	LogFileMaxSize    int    `json:"logFileMaxSize" yaml:"logFileMaxSize"`
	LogFileMaxBackups int    `json:"logFileMaxBackups" yaml:"logFileMaxBackups"`
	LogMaxAge         int    `json:"logMaxAge" yaml:"logMaxAge"`
}
