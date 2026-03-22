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
	name := flag.String("name", "Devops 學習", "輸入你的名字")
	owner := flag.String("owner", "yourusername", "GitHub 使用者名稱")
	repo := flag.String("repo", "devops-learning-journey", "repo 名稱")
	flag.Parse()

	fmt.Printf("Hello %s!, 這是我的 DevOps 學習旅程~\n", *name)
	fmt.Printf("今天是 %s，我成功用 ENTRYPOINT 跑起來了\n", time.Now().Format("2006-01-02"))

	// Call GitHub API get commit count
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
			fmt.Printf("JSON parase error, err: %v\n", err)
			return
		}
		fmt.Printf("你的 repo 目前有 %d 個 commit!\n", len(commits))
	} else {
		fmt.Println("(無法連上 GitHub... )")
		fmt.Printf("GitHub repsponse, err: %v", err)
	}
}
