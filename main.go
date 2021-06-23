package main

import (
	"github.com/Prithvipal/data-migration/constants"
	"github.com/Prithvipal/data-migration/processor"
)

func main() {
	processor := processor.NewProcessor("some/input", "some/output", constants.Csv, constants.Sql)
	processor.Process()
}
