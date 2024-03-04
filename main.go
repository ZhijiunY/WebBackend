package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	牛肉 = 10
	豬肉 = 7
	雞肉 = 5
)

// worker 模擬處理肉類的工作流程
// 創建一個名為 meats 的 channel，用於分配肉品給員工。
func worker(id string, meats <-chan string, wg *sync.WaitGroup) {
	// 當前 goroutine 完成後，通知 WaitGroup 減少一個計數
	defer wg.Done()
	// 循環讀取 meats 通道，直到此通道被關閉
	for meat := range meats {
		// 輸出員工開始處理肉類的日誌，包括員工ID、當前時間和肉類類型
		fmt.Printf("%s 在 %v 取得 %s\n", id, time.Now().Format("2006-01-02 15:04:05"), meat)
		// 定義一個變量來存儲處理時間
		var processTime time.Duration
		switch meat {
		// 根據肉類類型確定處理時間
		case "牛肉":
			processTime = 1 * time.Second
		case "豬肉":
			processTime = 2 * time.Second
		case "雞肉":
			processTime = 3 * time.Second
		}
		// 模擬處理肉類所需的時間
		time.Sleep(processTime)
		// 輸出員工完成處理肉類的日誌，包括員工ID、當前時間和肉類類型
		fmt.Printf("%s 在 %v 處理完 %s\n", id, time.Now().Format("2006-01-02 15:04:05"), meat)
	}
}

func main() {
	// 使用WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup
	// 創建一個有緩衝的通道，用於傳遞肉類信息
	meats := make(chan string, 22)

	// 創建五個員工的goroutine，每個員工對應一個 worker 函數的實例
	employees := []string{"A", "B", "C", "D", "E"}
	for _, e := range employees {
		// 為每個goroutine增加計數
		wg.Add(1)
		// 啟動goroutine，執行worker函數
		go worker(e, meats, &wg)
	}

	// 將肉類分別發送到meats通道中
	for i := 0; i < 10; i++ {
		meats <- "牛肉"
	}
	for i := 0; i < 7; i++ {
		meats <- "豬肉"
	}
	for i := 0; i < 5; i++ {
		meats <- "雞肉"
	}

	// 發送完畢後關閉通道
	close(meats)

	// 等待所有員工goroutine完成
	wg.Wait()
}
