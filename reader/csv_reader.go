package reader

type CsvReader struct {
	InputPath string
}

func (c CsvReader) Read(data []byte, offset int) (int, error) {
	return -1, nil
}
