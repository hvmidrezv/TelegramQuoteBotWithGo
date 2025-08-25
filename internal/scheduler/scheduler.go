package scheduler

import (
	"log"
	"telegrambot/internal/config"
	"telegrambot/internal/quotes"

	"github.com/go-co-op/gocron"
)

var (
	Scheduler *gocron.Scheduler
	jobs      = map[int64]*gocron.Job{}
)

type SendFunc func(chatID int64, text string) error

func Start(s *gocron.Scheduler, send SendFunc) {
	Scheduler = s
	s.StartAsync()
	ScheduleAll(send)
}

func ScheduleAll(send SendFunc) {
	for _, gid := range config.ConfigData.Groups {
		_ = ScheduleFor(gid, send)
	}
}

func ScheduleFor(chatID int64, send SendFunc) error {
	if job, ok := jobs[chatID]; ok {
		Scheduler.RemoveByReference(job)
		delete(jobs, chatID)
	}
	mins := config.IntervalFor(chatID)
	job, err := Scheduler.Every(mins).Minutes().Do(func() {
		q, err := quotes.FetchQuote()
		if err != nil {
			log.Println("fetchQuote err:", err)
			return
		}
		text := quotes.FormatQuote(q)
		_ = send(chatID, text)
	})
	if err != nil {
		return err
	}
	jobs[chatID] = job
	return nil
}
