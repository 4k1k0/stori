package domain

import "io"

type Reader interface {
	Read() (interface{}, error)
}

type ReaderIO interface {
	Read(r io.Reader) (interface{}, error)
}
