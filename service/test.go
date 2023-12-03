package service

import "time"

func Run() {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				startTime := time.Now()

				pml.service.LoadData(ctx)

				elapsedTime := time.Since(startTime)

				timeToNextTick := interval - elapsedTime

				if timeToNextTick < rest {
					holdTime := rest - timeToNextTick
					if timeToNextTick < 0 {
						holdTime = rest
					}

					time.Sleep(holdTime)
				}
			}
		}
	}()
}

func LoadData() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// go to database

	newMethods = make(map[string]YourEntity)

	for _, v := range yourData {
		newMethods[pmUuid] = v
	}

	pm.methods = newMethods
}

func FindCity(id int64) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	city, ok := pm.methods[id]
	if !ok {
		return // error
	}

	return city

}
