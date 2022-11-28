package main

import (
	"github.com/robfig/cron"
	"go-gin/models"
	"log"
	"time"
)

func main() {
	log.Println("Starting")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models Clean all tag")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models Clean all tag")
		models.CleanAllArticle()
	})

	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)

		}
	}
}
