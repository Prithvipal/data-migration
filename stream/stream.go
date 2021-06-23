package stream

type Streamer interface {
	Stream(data []interface{}) error
}
