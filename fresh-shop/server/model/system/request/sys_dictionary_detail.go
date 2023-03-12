package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
