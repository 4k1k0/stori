package domain

import "stori/pkg/result"

type Sender interface {
	Send(r result.Result) error
}
