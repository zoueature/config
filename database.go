package config

import (
	"fmt"
	"strings"
)

type DatabaseType string

const (
	mysqlType DatabaseType = "mysql"
	mongoType DatabaseType = "mongo"
)

type DatabaseConfig struct {
	Type     DatabaseType `json:"type" yaml:"type"`
	Username string       `json:"username" yaml:"username"`
	Password string       `json:"password" yaml:"password"`
	Database string       `json:"database" yaml:"database"`
	MinConn  *int         `json:"minConn" yaml:"minConn"`
	MaxConn  *int         `json:"maxConn" yaml:"maxConn"`

	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Charset  string `json:"charset" yaml:"charset"`
	Prefix   string `json:"prefix" yaml:"prefix"`
	Timezone string `json:"timezone" yaml:"timezone"`

	Nodes []string `json:"nodes" yaml:"nodes"`
}

func (d *DatabaseConfig) initDefaultInEmptyConfig() {
	if d.Charset == "" {
		d.Charset = defaultCharset
	}
	minConn := defaultMinConn
	if d.MinConn == nil {
		d.MinConn = &minConn
	}
	maxConn := defaultMaxConn
	if d.MaxConn == nil {
		d.MaxConn = &maxConn
	}
}

const (
	defaultCharset = "utf8mb4"
	defaultMinConn = 1
	defaultMaxConn = 100
)

func (d DatabaseConfig) Dsn() string {
	d.initDefaultInEmptyConfig()
	if d.Database == "" {
		panic("[Database Configuration] db name is empty")
	}
	s := ""
	switch d.Type {
	case mysqlType:
		s = d.mysqlDsn()
	case mongoType:
		s = d.toMongoDsn()
	default:
		panic(driverNotDefinedErr)
	}
	return s
}

func (d DatabaseConfig) mysqlDsn() string {
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

// mongodb://myDatabaseUser:D1fficultP%40ssw0rd@db0.example.com,db1.example.com,db2.example.com/?replicaSet=myRepl&w=majority&wtimeoutMS=5000
func (d DatabaseConfig) toMongoDsn() string {
	if len(d.Nodes) == 0 {
		panic("[Database Configuration] mongo node is empty. ")
	}
	auth := "%s:%s@"
	if d.Username == "" {
		auth = ""
	} else {
		auth = fmt.Sprintf(auth, d.Username, d.Password)
	}
	return fmt.Sprintf(
		"mongodb://%s%s/%s?minPoolSize=%d&maxPoolSize=%d",
		auth,
		strings.Join(d.Nodes, ","),
		d.Database,
		d.MinConn,
		d.MaxConn,
	)
}
