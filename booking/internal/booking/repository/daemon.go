package repository

import (
	"context"
	"log"
	"sync"
	"time"
)

var Jobs = make(chan int, 5)
var Result = make(chan string, 5)

func Daemon() {
	for {
		res := <-Result
		log.Printf("Worker does his work with result %s", res)
	}
}

func (r *Repo) Worker(wg *sync.WaitGroup, num int) {
	log.Printf("Worker number %d start work", num)
	wg.Done()
	for {
		id := <-Jobs
		log.Print("worker starts work")
		_, err := r.main.Exec("DELETE from bookrequest WHERE id=$1", id)
		if err != nil {
			log.Printf("worker cannot delete request with err: %v", err)
		}
		Result <- "ok"
	}
}

func (r *Repo) Hotels(mu *sync.RWMutex) {
	ticker := time.NewTicker(time.Minute)
	mu.Lock()
	hotels, err := r.GetHotels(context.TODO())
	if err != nil {
		log.Printf("err with get hotels: %v", err)
		return
	}
	mu.Unlock()
	RoomCounts := make(map[int]string, len(hotels))
	for {
		<-ticker.C
		mu.Lock()
		for _, hotel := range hotels {
			RoomCounts[hotel.Id] = hotel.Name
		}
		mu.Unlock()
	}
}
