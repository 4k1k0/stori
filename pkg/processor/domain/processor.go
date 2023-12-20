package domain

import "stori/pkg/result"

type Processor interface {
	Process() (result.Result, error)
}
