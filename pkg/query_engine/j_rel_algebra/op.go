package rel_algebra

import (
	"bytes"
	process "colexecdb/pkg/query_engine/e_process"
	"colexecdb/pkg/query_engine/j_rel_algebra/projection"
)

var stringFunc = [...]func(any, *bytes.Buffer){
	Projection: projection.String,
}

var prepareFunc = [...]func(*process.Process, any) error{
	Projection: projection.Prepare,
}

var execFunc = [...]func(*process.Process, any) (process.ExecStatus, error){
	Projection: projection.Call,
}
