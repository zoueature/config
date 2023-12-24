package config

type StorageConfig struct {
	AccessKey    string `json:"accessKey" yaml:"accessKey"`
	AccessSecret string `json:"accessSecret" yaml:"accessSecret"`
	Bucket       string `json:"bucket" yaml:"bucket"`
	Region       string `json:"region" yaml:"region"`
	Domain       string `json:"domain" yaml:"domain"`
	CDN          string `json:"cdn" yaml:"cdn"`
}
