package domain

type Processor interface {
	Process() (interface{}, error)
}
