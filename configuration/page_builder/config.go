package page_builder

import "reflect"

const (
	NumericInputType = "number"
	StringInputType  = "text"
	DefaultInputType = "text"
)

var (
	fieldKindToInputType = map[string]string{
		reflect.Int.String():     NumericInputType,
		reflect.Int32.String():   NumericInputType,
		reflect.Int64.String():   NumericInputType,
		reflect.Uint.String():    NumericInputType,
		reflect.Uint32.String():  NumericInputType,
		reflect.Uint64.String():  NumericInputType,
		reflect.Float64.String(): NumericInputType,
		reflect.Float32.String(): NumericInputType,
		reflect.String.String():  StringInputType,
	}
)
