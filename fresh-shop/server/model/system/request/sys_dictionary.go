package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
