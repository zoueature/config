package config

import (
	"errors"
	"sync"
)
import (
	"github.com/spf13/cast"
)

const (
	defaultListen = ":8000"
)

var (
	driverNotDefinedErr = errors.New("driver not defined")

	defaultLogConfig = &LogConfig{
		Driver:   "stdout",
		MinLevel: "debug",
	}
	defaultAppConfig = &AppConfig{
		Listen: defaultListen,
		Debug:  true,
	}

	defaultConfig = &Configuration{
		App:       defaultAppConfig,
		Log:       defaultLogConfig,
		otherData: sync.Map{},
	}
)

func DefaultConfig() *Configuration {
	return defaultConfig
}

func DefaultAppConfig() *AppConfig {
	return defaultAppConfig
}

func DefaultLogConfig() *LogConfig {
	return defaultLogConfig
}

type Conf interface {
	AppConfig() *AppConfig
	LogConfig() *LogConfig
	SetAppConfig(config *AppConfig)
	SetLogConfig(config *LogConfig)
}

type Configuration struct {
	App       *AppConfig      `json:"app" yaml:"app"`
	Log       *LogConfig      `json:"log" yaml:"log"`
	Database  *DatabaseConfig `json:"database" yaml:"database"`
	Cache     *CacheConfig    `json:"cache" yaml:"cache"`
	otherData sync.Map
}

func (c *Configuration) AppConfig() *AppConfig {
	return c.App
}
func (c *Configuration) LogConfig() *LogConfig {
	return c.Log
}
func (c *Configuration) SetAppConfig(config *AppConfig) {
	c.App = config
}
func (c *Configuration) SetLogConfig(config *LogConfig) {
	c.Log = config
}

func (c *Configuration) Set(key string, value interface{}) {
	c.otherData.Store(key, value)
}

func (c *Configuration) Get(key string) (interface{}, bool) {
	return c.otherData.Load(key)
}

func (c *Configuration) MustGet(key string) interface{} {
	v, ok := c.otherData.Load(key)
	if !ok {
		panic("configuration: " + key + " not exists")
	}
	return v
}

func (c *Configuration) MustGetString(key string) string {
	v := c.MustGet(key)
	return cast.ToString(v)
}

func (c *Configuration) ShouldGetString(key string) (string, bool) {
	v, ok := c.Get(key)
	if !ok {
		return "", false
	}
	return cast.ToString(v), ok
}

func (c *Configuration) MustGetInt64(key string) int64 {
	v := c.MustGet(key)
	return cast.ToInt64(v)
}

func (c *Configuration) ShouldGetInt64(key string) (int64, bool) {
	v, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	return cast.ToInt64(v), ok
}

func (c *Configuration) MustGetInt(key string) int {
	v := c.MustGet(key)
	return cast.ToInt(v)
}

func (c *Configuration) ShouldGetInt(key string) (int, bool) {
	v, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	return cast.ToInt(v), ok
}
