package main

import (
	"context"
	"fmt"
	"github.com/matheus-gondim/hydration-reminder/config"
	"github.com/matheus-gondim/hydration-reminder/internal/domain"
	"github.com/matheus-gondim/hydration-reminder/internal/infra/notifier"
	"log"
	"time"
)

func main() {
	cfg := config.Envs
	now := time.Now()

	user := domain.User{
		Weight:        cfg.Weight,
		LunchInterval: float64(cfg.LunchIntervalMinutes) / 60,
		OfficeHours:   cfg.OfficeHours,
	}

	inactiveStart := time.Date(now.Year(), now.Month(), now.Day(), cfg.LunchIntervalStart, 0, 0, 0, time.Local)
	inactiveEnd := inactiveStart.Add(time.Duration(cfg.LunchIntervalMinutes) * time.Minute)

	activeEnd := now.Add(time.Duration(cfg.OfficeHours) * time.Hour)
	latestEnd := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)
	if activeEnd.After(latestEnd) {
		activeEnd = latestEnd
	}

	scheduler := domain.NewScheduler(now, activeEnd, inactiveStart, inactiveEnd)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go scheduler.HandleInterrupt(cancel)
	scheduler.Start(ctx, user.DailyWaterIntakeInGlassesPerOfficeHours(200), func() {
		err := notifier.Notify("Beba água", fmt.Sprintf("Beba %vml de água", 200))
		if err != nil {
			log.Printf("Erro ao exibir notificação: %v", err)
		}
	})
	<-ctx.Done()
}
