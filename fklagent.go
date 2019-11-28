package main

import (
	_ "fasoo.com/fklagent/db"
	"fasoo.com/fklagent/process"
	"fasoo.com/fklagent/util/log"
	"sync"
)

func main() {
	log.INFO("start processing data!!")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		process.SFTP()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		process.ESB()
	}()

	wg.Wait()
}
