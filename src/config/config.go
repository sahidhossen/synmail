package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                     int    `envconfig:"PORT" required:"true" default:"8080"`
	Secret                   string `envconfig:"SECRET" required:"true"`
	DebugMode                bool   `envconfig:"DEBUG_MODE" required:"true"`
	DatabaseConnectionString string `envconfig:"DATABASE_URL" required:"true"`
	BasePath                 string
	MailMailer               string `envconfig:"MAIL_MAILER" required:"true"`
	MailHost                 string `envconfig:"MAIL_HOST" required:"true"`
	MailPort                 int    `envconfig:"MAIL_PORT" required:"true" default:"587"`
	MailUsername             string `envconfig:"MAIL_USERNAME" required:"true"`
	MailPass                 string `envconfig:"MAIL_PASSWORD" required:"true"`
	MailFrom                 string `envconfig:"MAIL_FROM_ADDRESS" required:"true"`
	MailFromName             string `envconfig:"MAIL_FROM_NAME" required:"true"`
	MailAPIKey               string `envconfig:"MAIL_API_KEY"`
	MailGunDomain            string `envconfig:"MAIL_GUN_DOMAIN"`
	AWSRegion                string `envconfig:"MAIL_AWS_REGION"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	cfg.BasePath = getStoragePath()
	return &cfg, err
}

func getStoragePath() string {
	path, _ := os.Getwd()
	storagePath := filepath.Join(filepath.Dir(path), "storage")
	err := os.MkdirAll(storagePath, os.ModePerm)
	if err != nil {
		fmt.Printf("Fail to create storage folder")
		return ""
	}
	return storagePath
}
