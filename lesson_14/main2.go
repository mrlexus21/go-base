package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		resultCh    = make(chan string)
		ctx, cancel = context.WithCancel(context.Background())
		services    = []string{"Super", "Villagemobil", "Sett Taxi", "Index Go"}
		wg          sync.WaitGroup
		winner      string
	)

	defer cancel()

	for i := range services {
		svc := services[i]

		wg.Add(1)
		go func(serviceName string) {
			defer wg.Done()
			requestRide(ctx, serviceName, resultCh)
		}(svc)
	}

	go func() {
		winner = <-resultCh
		cancel()
		close(resultCh) // Закрываем канал после получения результата
	}()

	wg.Wait()
	log.Printf("found car in %q", winner)
}

func requestRide(ctx context.Context, serviceName string, resultCh chan string) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			time.Sleep(3 * time.Second) // Перемещаем sleep внутрь цикла
			if rand.Float64() > 0.75 {
				select {
				case resultCh <- serviceName:
				case <-ctx.Done():
					return
				}
				return
			}
		}
	}
}
