package config

import (
	"database/sql"
	"embed"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var config *Cfg
var once sync.Once

type Cfg struct {
	Database *sql.DB
	FS       embed.FS
	FileName string
	Email    string
	S3Config S3Config
}

type S3Config struct {
	Bucket   string
	Key      string
	S3Client *s3.Client
}

// New sets the project config values using a singleton.
func New(cfg Cfg) *Cfg {
	once.Do(func() {
		config = &Cfg{
			Database: cfg.Database,
			FS:       cfg.FS,
			FileName: cfg.FileName,
			Email:    cfg.Email,
			S3Config: cfg.S3Config,
		}
	})
	return config
}

// Config retrieves the project's config.
func Config() *Cfg {
	return config
}
