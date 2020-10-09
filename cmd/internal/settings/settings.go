package settings

import (
	"github.com/Netflix/go-env"
)

func New() (*Settings, error) {
	settings := &Settings{}
	_, err := env.UnmarshalFromEnviron(settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

type Settings struct {
	AWSSettings
	LogLevel string `env:"LOG_LEVEL,default=info"`
	Debug    bool   `env:"DEBUG,default=false"`
}

type AWSSettings struct {
	ID       string `env:"AWS_ACCESS_KEY_ID"`
	Secret   string `env:"AWS_ACCESS_KEY_SECRET"`
	Region   string `env:"AWS_REGION"`
	Endpoint string `env:"AWS_ENDPOINT"`
	Bucket   string `env:"S3_BUCKET"`
}
