package config

import (
	"embed"
	"sync"
)

var config *Cfg
var once sync.Once

type Cfg struct {
	FS       embed.FS
	FileName string
	Email    string
}

// New sets the project config values using a singleton.
func New(fs embed.FS, filename, email string) *Cfg {
	once.Do(func() {
		config = &Cfg{
			FS:       fs,
			FileName: filename,
			Email:    email,
		}
	})
	return config
}

// Config retrieves the project's config.
func Config() *Cfg {
	return config
}
