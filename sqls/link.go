package sqls

import (
	_ "unsafe"
)

//go:linkname convertAssign database/sql.convertAssign
//go:nosplit
func convertAssign(dest, src any) error

func ConvertAssign[T any](dest T, src any) error { return convertAssign(dest, src) }
