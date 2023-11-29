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
}

func (d DatabaseConfig) Dsn() (string, error) {
	switch d.Type {
	case mysqlType:
		return d.ToMySQLDsn(), nil
	default:
		return "", driverNotDefinedErr
	}
}

func (d DatabaseConfig) ToMySQLDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
		d.Charset,
	)
}
