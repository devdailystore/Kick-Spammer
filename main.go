package main

import (
	"fmt"
	"kick-spammer/utils"
	"sync"
	"time"
)

func worker(channelId int, message string, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		proxy := utils.GetRandomProxy()
		token := utils.GetRandomToken()

		if proxy == "" || token == "" {
			fmt.Printf("Worker %d: No proxies or tokens available, retrying...\n", workerID)
			time.Sleep(100 * time.Millisecond)
			return
		}

		fmt.Printf("Worker %d using proxy: %s\n", workerID, proxy)
		err := sendMessage(channelId, message, proxy, token)
		if err != nil {
			fmt.Printf("Worker %d: Error - %v (continuing...)\n", workerID, err)
		}
		if err == nil {
			fmt.Printf("Worker %d: Message sent successfully\n", workerID)
		}
	}
}

func main() {
	numThreads := 100
	channelName := "lordrobson"
	message := "cwel"

	userInfo := utils.GetUserId(channelName)
	if userInfo.UserId == -1 {
		fmt.Printf("Failed to get user ID for channel '%s'. Exiting...\n", channelName)
		return
	}

	channelId := userInfo.UserId
	if !userInfo.CanSpam {
		fmt.Printf("Channel '%s' has followers mode enabled. Cannot spam\n", channelName)
		return
	}

	fmt.Printf("Channel ID for '%s': %d\n", channelName, channelId)
	fmt.Printf("Starting %d worker threads...\n", numThreads)

	var wg sync.WaitGroup

	for i := range numThreads {
		wg.Add(1)
		go worker(channelId, message, i+1, &wg)
	}

	wg.Wait()
}
