package main

import (
	"github.com/robfig/cron"
	"go-gin-example/models"
	"log"
	"time"
)

func main() {
	log.Println("Starting......")

	// 创建一个新的cron job runner
	c := cron.New()

	// AddFunc 向cron job runner 添加一个func 然后按照时间表运行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
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
