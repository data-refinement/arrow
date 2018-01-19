package arrow

import (
	"github.com/influxdata/arrow/memory"
)

// ArrayData is a self-contained representation of the memory and metadata
// which represents an Arrow Array.
type ArrayData struct {
	typE      DataType
	nullN     int
	length    int
	buffers   []*memory.Buffer // TODO(sgc): should this be an interface?
	childData []*ArrayData     // TODO(sgc): used by ListArray, StructArray and UnionArray
}

func NewArrayData(typE DataType, length int, buffers []*memory.Buffer, nullN int) *ArrayData {
	return &ArrayData{
		typE:    typE,
		nullN:   nullN,
		length:  length,
		buffers: buffers,
	}
}