package config

import (
	"io"
	"io/fs"
	"os"
)

const (
	StdDriver  stdDriver  = "stdout"
	FileDriver fileDriver = "file"
)

type LogDriver interface {
	GetWriter() (io.Writer, error)
	String() string
}

type stdDriver string

func (d stdDriver) GetWriter() (io.Writer, error) {
	return os.Stdout, nil
}

func (d stdDriver) String() string {
	return string(d)
}

type fileDriver string

func (d fileDriver) GetWriter() (io.Writer, error) {
	fd, err := os.OpenFile(string(d), os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	return fd, nil
}

func (d fileDriver) String() string {
	return string(d)
}

type LogConfig struct {
	DriverType string `json:"driver" yaml:"driver"`
	FilePath   string `json:"filePath" yaml:"filePath"`
	driver     LogDriver
}

// generateDriverBy 生成driver
func (c *LogConfig) generateDriverBy() error {
	switch c.DriverType {
	case string(StdDriver):
		c.driver = StdDriver
	case string(FileDriver):
		c.driver = FileDriver
	default:
		return driverNotDefinedErr
	}
	return nil
}

func (c *LogConfig) GetLogWriter() (io.Writer, error) {
	if c.driver == nil {
		err := c.generateDriverBy()
		if err != nil {
			return nil, err
		}
	}
	return c.driver.GetWriter()
}

func (c *LogConfig) SetDriver(driver LogDriver) {
	c.driver = driver
}
