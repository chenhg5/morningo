package schedule

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
	"os"
	"morningo/config"
	"morningo/modules/logger"
)

func init() {
	c := cron.New()

	// 切割日志
	c.AddFunc("0 59 23 * * *", func() {
		cutLog(config.GetEnv().ACCESS_LOG_PATH)
		cutLog(config.GetEnv().INFO_LOG_PATH)
		cutLog(config.GetEnv().ERROR_LOG_PATH)
		logger.InitAllLogger()
	})

	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.AddFunc("@every 5s", func() { fmt.Println("Every five seconds") })
	c.Start()
}

func cutLog(path string) {
	date := time.Now().Format("20060102")
	err := os.Rename(path, path+"."+date+".log")
	if err != nil {
		fmt.Println(path + " rename Error!")
	} else {
		fmt.Println(path + " rename OK!")
	}
}