package config

import (
    "fmt"

  "github.com/robfig/cron"
)

struct Cron {
  

func InitCron() {
  c := *cron.New()

  c.AddFunc("0 * * * *", func() { fmt.Println("Every minute") })
  c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
  c.AddFunc("@every 1h30m", func() { fmt.Println("Every 1 hour 30 minutes") })
  c.AddFunc("@daily", func() { fmt.Println("Every day") })

  c.Start()

  // Mantém o programa em execução para permitir a execução das funções agendadas
  select {}
}
