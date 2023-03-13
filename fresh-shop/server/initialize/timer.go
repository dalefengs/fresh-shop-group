package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"fresh-shop/server/config"
	"fresh-shop/server/global"
	"fresh-shop/server/utils"
)

func Timer() {
	if global.Config.Timer.Start {
		for i := range global.Config.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.Config.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.Timer.AddTaskByFunc("ClearDB", global.Config.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.Config.Timer.Detail[i])
		}
	}
}
