package convert

type Converter interface {
	Convert(data []byte) []interface{}
}
