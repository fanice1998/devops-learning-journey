package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	name := flag.String("name", "DevOps 學習者", "輸入你的名字")
	owner := flag.String("owner", "yourusername", "GitHub 使用者名稱")
	repo := flag.String("repo", "devops-learning-journey", "repo 名稱")
	flag.Parse()

	fmt.Printf("Hello %s\n", *name)
	fmt.Printf("今天日期：%s\n", time.Now().Format("2006-01-02"))

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", *owner, *repo)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "my-devops-tool")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var commits []interface{}
		err = json.Unmarshal(body, &commits)
		if err != nil {
			fmt.Printf("JSON 解析錯誤: %v\n", err)
			return
		}
		fmt.Printf("目前 commit 數量：%d\n", len(commits))
	} else {
		fmt.Println("無法連上 GitHub")
		if err != nil {
			fmt.Printf("錯誤: %v\n", err)
		}
	}
}
