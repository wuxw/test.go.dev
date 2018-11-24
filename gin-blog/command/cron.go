package command

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"test.go.dev/gin-blog/models"
	"test.go.dev/gin-blog/pkg/logging"
	"test.go.dev/gin-blog/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	log.Println("Starting...")
	c := cron.New()
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
