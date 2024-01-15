package main

import (
	"dmp-training/configs"
	"dmp-training/src"
	"sync"
	"time"
)

var cfg configs.Config

func init() {
	cfg = configs.NewConfig()
}
func main() {
	server := src.NewServer(cfg)
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Local = time.FixedZone("UTC+7", 7*60*60)

		server.Run()
	}()

	wg.Wait()

}
