package transform

type Transformer interface {
	Transform(data []interface{}) []interface{}
}

type DataTranformer struct {
}

func (d DataTranformer) Transform(data []interface{}) []interface{} {
	return nil
}
