package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	imageURL := "https://i.imgur.com/5mpS7Ij.jpg"
	fileName := "5mpS7Ij.jpg"

	// 建立一個新的請求
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		fmt.Printf("無法建立請求: %v\n", err)
		return
	}

	// 設定 User-Agent 標頭，模擬瀏覽器行為
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// 發送請求
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("無法下載圖片: %v\n", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("下載失敗，狀態碼: %d\n", response.StatusCode)
		return
	}

	// 創建一個新檔案來保存圖片
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("無法創建檔案: %v\n", err)
		return
	}
	defer file.Close()

	// 將響應內容複製到檔案中
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("無法保存圖片: %v\n", err)
		return
	}

	fmt.Printf("圖片已成功下載並保存為 %s\n", fileName)
}
