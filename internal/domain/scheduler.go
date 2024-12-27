package domain

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

type Period struct {
	Start time.Time
	End   time.Time
}

type Scheduler struct {
	ActivePeriod   Period
	InactivePeriod Period
}

func NewScheduler(activeStart, activeEnd, inactiveStart, inactiveEnd time.Time) *Scheduler {
	return &Scheduler{
		ActivePeriod:   Period{Start: activeStart, End: activeEnd},
		InactivePeriod: Period{Start: inactiveStart, End: inactiveEnd},
	}
}

func (s *Scheduler) Start(ctx context.Context, frequency int, action func()) {
	if frequency <= 0 {
		log.Println("A frequência deve ser maior que zero.")
		return
	}

	interval := time.Hour / time.Duration(frequency)
	ticker := time.NewTicker(interval)

	log.Printf("Scheduler iniciado: intervalo de %v\n", interval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("Contexto cancelado. Encerrando o Scheduler...")
				return
			case t := <-ticker.C:
				if !s.isWithinInterval(t) {
					log.Printf("Fora do intervalo ativo ou dentro do intervalo inativo: %v\n", t)
					continue
				}
				
				action()
				log.Printf("Executando ação às %v\n", t)
			}
		}
	}()
}

func (s *Scheduler) isWithinInterval(t time.Time) bool {
	isActive := t.After(s.ActivePeriod.Start) && t.Before(s.ActivePeriod.End)
	isInactive := t.After(s.InactivePeriod.Start) && t.Before(s.InactivePeriod.End)
	return isActive && !isInactive
}

func (s *Scheduler) HandleInterrupt(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Sinal de interrupção recebido, encerrando...")
	cancel()
}
