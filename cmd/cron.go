package cmd

import (
    "fmt"

  "github.com/robfig/cron"
)

type Entry struct{
  Yearly string
  Monthly string
  Weekly string
  Daily string
  Hourly string
  Minutely string
}

var entries = Entry{
  Yearly: "@yearly",
  Monthly: "@monthly",
  Weekly: "@weekly",
  Daily: "@daily",
  Hourly: "@hourly",
  Minutely: "@minutely",
}

func InitCron() {
  c := *cron.New()

  c.AddFunc("0 * * * *", func() { fmt.Println("Every minute") })
  c.AddFunc(entries.Minutely, func() { fmt.Println("Every hour") })
  c.AddFunc("@every 1h30m", func() { fmt.Println("Every 1 hour 30 minutes") })
  c.AddFunc(entries.Daily, func() { fmt.Println("Every day") })

  c.Start()

  // Keep the program running to allow scheduled functionsto execute
  select {}
}

func Stop(c *cron.Cron) {
  c.Stop()
}
