package reader

type Reader interface {
	Read(data []byte, offset int) (int, error)
}
