package processor

import (
	"github.com/Prithvipal/data-migration/constants"
	"github.com/Prithvipal/data-migration/convert"
	"github.com/Prithvipal/data-migration/reader"
	"github.com/Prithvipal/data-migration/stream"
	"github.com/Prithvipal/data-migration/transform"
)

func NewProcessor(inputPath string, outputPath string, input constants.InputType, output constants.OutputType) dataProcessor {
	return dataProcessor{
		reader:      reader.CsvReader{InputPath: inputPath},
		converter:   &convert.CsvConverter{},
		transformer: transform.DataTranformer{},
		streamer:    stream.SqlStreamer{OutputPath: outputPath},
	}
}

type dataProcessor struct {
	reader      reader.Reader
	converter   convert.Converter
	transformer transform.Transformer
	streamer    stream.Streamer
}

func (d dataProcessor) Process() {
	var data []byte
	offset := 0
	for i, err := d.reader.Read(data, offset); i != -1 && err != nil; {
		offset = offset + i
		davaValue := d.converter.Convert(data)
		davaValueChanged := d.transformer.Transform(davaValue)
		go d.streamer.Stream(davaValueChanged)
	}

}
