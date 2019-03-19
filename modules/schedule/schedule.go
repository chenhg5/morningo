package schedule

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
	"os"
	"morningo/config"
	"morningo/modules/log"
)

func init() {
	c := cron.New()

	// 切割日志
	c.AddFunc("0 59 23 * * *", func() {
		cutLog(config.GetEnv().ACCESS_LOG_PATH)
		cutLog(config.GetEnv().INFO_LOG_PATH)
		cutLog(config.GetEnv().ERROR_LOG_PATH)
		log.InitAllLogger()
	})

	c.AddFunc("0 30 * * * *", func() { log.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { log.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { log.Println("Every hour thirty") })
	c.AddFunc("@every 5s", func() { log.Println("Every five seconds") })
	c.Start()
}

func cutLog(path string) {
	date := time.Now().Format("20060102")
	err := os.Rename(path, path+"."+date+".log")
	if err != nil {
		log.Println(path + " rename Error!")
	} else {
		log.Println(path + " rename OK!")
	}
}