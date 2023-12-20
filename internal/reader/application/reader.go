package application

import (
	"stori/internal/reader/domain"
	"stori/internal/reader/infrastructure"
)

func New() domain.Reader {
	return infrastructure.New("")
}
