package main

import (
	"sync"
	"time"

	"github.com/muflikhandimasd/golang-basic-clean/configs"
	"github.com/muflikhandimasd/golang-basic-clean/src"
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
