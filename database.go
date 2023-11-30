package config

import "fmt"

type DatabaseType string

const (
	mysqlType DatabaseType = "mysql"
	mongoType DatabaseType = "mongo"
)

type DatabaseConfig struct {
	Type     DatabaseType `json:"type" yaml:"type"`
	Host     string       `json:"host" yaml:"host"`
	Port     string       `json:"port" yaml:"port"`
	Username string       `json:"username" yaml:"username"`
	Password string       `json:"password" yaml:"password"`
	Database string       `json:"database" yaml:"database"`
	Charset  string       `json:"charset" yaml:"charset"`
	Prefix   string       `json:"prefix" yaml:"prefix"`
	Timezone string       `json:"timezone" yaml:"timezone"`
}

const defaultCharset = "utf8mb4"

func (d DatabaseConfig) Dsn() string {
	if d.Charset == "" {
		d.Charset = defaultCharset
	}
	if d.Database == "" {
		panic("[Database Configuration] db name is empty")
	}
	s := ""
	switch d.Type {
	case mysqlType:
		s = d.ToMySQLDsn()
	default:
		panic(driverNotDefinedErr)
	}
	return s
}

func (d DatabaseConfig) ToMySQLDsn() string {
	if d.Timezone == "" {
		d.Timezone = "Local"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.Charset,
		d.Timezone,
	)
}
