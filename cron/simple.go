package cron

import (
	"fmt"

	"github.com/ThreeKing2018/goutil/choke"
	"github.com/robfig/cron"
	"github.com/yezihack/gotime"
)

func AddCron() {
	c := cron.New()
	c.Start()
	defer c.Stop()
	c.AddFunc("* * * * * *", func() {
		fmt.Println(gotime.New().Now())
	})
	c.AddFunc("1/5 * * * * *", func() {
		fmt.Println("5秒一次", gotime.New().Now())
	})
	c.AddFunc("* 1 * * * *", func() {
		fmt.Println("----------------------每小时里的第一分钟都执行1:00-1:59,如12:01:01-12:01:59都执行", gotime.New().Now())
	})
	c.AddFunc("1 */1 * * * *", func() {
		fmt.Println("----------------------每分钟01秒执行一次,如12:45:01", gotime.New().Now())
	})
	c.AddFunc("@every 10s", func() {
		fmt.Println("----------------------每10秒执行", gotime.New().Now())
	})
	c.AddFunc("@every 1m", func() {
		fmt.Println("----------------------每1分钟执行", gotime.New().Now())
	})
	//阻塞
	choke.Choke()
}
