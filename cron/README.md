# cron go 定时器
* Cron表达式表示一组时间，使用 6 个空格分隔的字段
* Golang 的 Cron 比 Crontab 多了一个秒级，以后遇到秒级要求的时候就省事了

# 安装
```
go get -u github.com/robfig/cron
```

字段名   | 是否必填 | 允许的值  | 允许的特殊字符
----------   | ---------- | --------------  | --------------------------
秒（Seconds）      | Yes        | 0-59            | * / , -
分（Minutes）      | Yes        | 0-59            | * / , -
时（Hours）        | Yes        | 0-23            | * / , -
一个月中的某天（Day of month） | Yes        | 1-31            | * / , - ?
月（Month）        | Yes        | 1-12 or JAN-DEC | * / , -
星期几（Day of week）  | Yes        | 0-6 or SUN-SAT  | * / , - ?


### 预定义的 Cron 时间表
输入                  | 简述                                | 相当于
-----                  | -----------                                | -------------
@yearly (or @annually) | 1月1日午夜运行一次      | 0 0 0 1 1 *
@monthly               | 每个月的午夜，每个月的第一个月运行一次 | 0 0 0 1 * *
@weekly                | 每周一次，周日午夜运行一次       | 0 0 0 * * 0
@daily (or @midnight)  | 每天午夜运行一次                  | 0 0 0 * * *
@hourly                | 每小时运行一次        | 0 0 * * * *

# 简单使用
```
c := cron.New()
c.Start()
defer  c.Stop()
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

```

# 实例
- [实例](simple.go)