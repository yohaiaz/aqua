package services

import (
	"context"
	"encoding/json"
	"myrepo/aqua/entities"
	"myrepo/aqua/sender"
	"sync"
)

type Crawler struct {
	canceller               context.CancelFunc
	wg                      sync.WaitGroup
	semAllConcurrentSenders chan bool
	filesChan               chan *entities.File
	senderApi               sender.ISender
}

func NewCrawler(files chan *entities.File, sender sender.ISender) *Crawler {

	return &Crawler{
		wg: sync.WaitGroup{},
		semAllConcurrentSenders: make(chan bool, 3),
		filesChan:               files,
		senderApi:               sender,
	}
}

func (c *Crawler) Crawl() {
	for f := range c.filesChan {
		c.semAllConcurrentSenders <- true
		c.wg.Add(1)

		go func(file *entities.File) {
			defer func() {
				<-c.semAllConcurrentSenders
				c.wg.Done()
			}()

			b, _ := json.Marshal(file)

			c.senderApi.Send(string(b))
		}(f)
	}

	c.wg.Wait()
}