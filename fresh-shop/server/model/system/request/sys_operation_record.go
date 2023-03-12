package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
