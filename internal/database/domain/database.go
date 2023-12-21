package domain

import "database/sql"

type Connector interface {
	Connect() (*sql.DB, error)
	Disconnect() error
}
