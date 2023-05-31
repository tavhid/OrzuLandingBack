package config

import (
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(ImportConfigs)

// Postgres configuration
type Postgres struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	SSLMode      string `mapstructure:"ssl_mode"`
	DatabaseName string `mapstructure:"database_name"`
}

// Redis configuration
type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBIndex  int    `json:"db_index"`
	Password string `json:"password"`
}

// Server configuration
type Server struct {
	Host      string `json:"host"`
	Port      uint16 `json:"port"`
	SecretKey string `mapstructure:"secret_key"`
}

// SMS configuration
type SMS struct {
	User              string `json:"user"`
	Password          string `json:"password"`
	URL               string `json:"url"`
	From              string `json:"from"`
	RegistrationText  string `mapstructure:"registration_text"`
	ResetPasswordText string `mapstructure:"reset_password_text"`
}

// Config ...
type Config struct {
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
	Server   Server   `json:"server"`
	SMS      SMS      `mapstructure:"sms_service"`
}
