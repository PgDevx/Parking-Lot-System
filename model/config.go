package model

import "time"

type Config struct {
	ServerConfig   ServerConfig   `mapstructure:"server"`
	DatabaseConfig DatabaseConfig `mapstructure:"database"`
	LoggerConfig   LoggerConfig   `mapstructure:"logger"`
	CORSConfig     CORSConfig     `mapstructure:"cors"`
}

type ServerConfig struct {
	Env          string        `mapstructure:"env"`
	Port         string        `mapstructure:"port"`
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	CloseTimeout time.Duration `mapstructure:"closeTimeout"`
	CSRFProtect  string        `mapstructure:"csrf_protect"`
	StaticDir    string        `mapstructure:"static_dir"`
}

// DatabaseConfig has database related configuration.
type DatabaseConfig struct {
	Type             string `mapstructure:"type"`
	Host             string `mapstructure:"host"`
	DbName           string `mapstructure:"dbName"`
	UserName         string `mapstructure:"userName"`
	Password         string `mapstructure:"password"`
	ConnectionString string `mapstructure:"connectionString"`
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string `mapstructure:"file"`
}

type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
}
