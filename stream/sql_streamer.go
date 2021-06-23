package stream

type SqlStreamer struct {
	OutputPath string
}

func (s SqlStreamer) Stream(data []interface{}) error {
	return nil
}
