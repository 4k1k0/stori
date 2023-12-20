package infrastructure

type ReaderCSV struct {
	FilePath string
}

func New(file string) *ReaderCSV {
	return &ReaderCSV{
		FilePath: file,
	}
}

func (r *ReaderCSV) Read() (interface{}, error) {
	return nil, nil
}
