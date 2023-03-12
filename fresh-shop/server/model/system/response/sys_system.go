package response

import "fresh-shop/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
