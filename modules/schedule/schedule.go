package schedule

import (
	"github.com/robfig/cron"
	"morningo/config"
	"morningo/modules/log"
	"os"
	"time"
)

func init() {
	c := cron.New()

	// 切割日志
	_ = c.AddFunc("0 59 23 * * *", func() {
		cutLog(config.GetEnv().AccessLogPath)
		cutLog(config.GetEnv().InfoLogPath)
		cutLog(config.GetEnv().ErrorLogPath)
		log.InitAllLogger()
	})

	_ = c.AddFunc("0 30 * * * *", func() { log.Println("Every hour on the half hour") })
	_ = c.AddFunc("@hourly", func() { log.Println("Every hour") })
	_ = c.AddFunc("@every 1h30m", func() { log.Println("Every hour thirty") })
	_ = c.AddFunc("@every 5s", func() { log.Println("Every five seconds") })
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
