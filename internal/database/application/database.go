package application

import (
	"stori/internal/database/domain"
	"stori/internal/database/infrastructure"
)

func New() domain.Connector {
	return &infrastructure.PostgresConnector{}
}
