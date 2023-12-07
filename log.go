package config

const (
	StdoutDriver = "stdout"
	FileDriver   = "file"
)

type DingtalkNotifySecret struct {
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	SignSecret  string `json:"signSecret" yaml:"signSecret"`
}

type LogConfig struct {
	Driver            string                `json:"driver" yaml:"driver"`
	Path              string                `json:"path" yaml:"path"`
	MinLevel          string                `json:"minLevel"`
	LogFileMaxSize    int                   `json:"logFileMaxSize" yaml:"logFileMaxSize"`
	LogFileMaxBackups int                   `json:"logFileMaxBackups" yaml:"logFileMaxBackups"`
	LogMaxAge         int                   `json:"logMaxAge" yaml:"logMaxAge"`
	DingtalkAlarm     *DingtalkNotifySecret `json:"dingtalkAlarm" yaml:"dingtalkAlarm"`
}
