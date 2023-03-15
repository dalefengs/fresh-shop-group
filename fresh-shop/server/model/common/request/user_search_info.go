package request

// UserSearch 用户信息搜索结构体
type UserSearch struct {
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
}
