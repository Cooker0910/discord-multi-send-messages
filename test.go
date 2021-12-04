package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/fatih/color"
)

type jsonResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ExitSafely() {
	color.Red("\nPress Enter to EXIT")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	color.Blue("\r\n\r\n\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2584    \u2584\u2584\u2584\u2584\u2588\u2588\u2588\u2584\u2584\u2584\u2584   \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2584          \u2584\u2588\u2588\u2588\u2588\u2588\u2588\u2584   \u2584\u2588\u2588\u2588\u2588\u2588\u2588\u2584  \r\n\u2588\u2588\u2588   \u2580\u2588\u2588\u2588 \u2584\u2588\u2588\u2580\u2580\u2580\u2588\u2588\u2588\u2580\u2580\u2580\u2588\u2588\u2584 \u2588\u2588\u2588   \u2580\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2580  \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588       \u2584\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588      \u2580\u2580\u2588\u2588\u2588 \u2588\u2588\u2588\u2588\u2584  \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588   \u2584\u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588   \u2584\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580   \u2580\u2588   \u2588\u2588\u2588   \u2588\u2580  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580         \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580   \u2580\u2588\u2588\u2588\u2588\u2588\u2588\u2580  \r\n                                                                   \r\n\rDISCORD MASS DM GO for RetroByte's NFT Launch")
	time.Sleep(2 * time.Second)
	color.White("Menu:\n1) Mass DM advertiser\n2) Single DM spam3) Token Checker\n4) Exit")
	color.White("\nEnter your choice: ")
	var choice int
	fmt.Scanln(&choice)
	if choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 5 && choice != 6 && choice != 7 && choice != 8 && choice != 9 && choice != 0 {
		color.Red("[%v] Invalid choice", time.Now().Format("15:05:04"))
		return
	}
	switch choice {
	case 1:
		var completed []string
		var failed []string
		var dead []string
		tokens, err := ReadLine("tokens.txt")
		members, err := ReadLine("memberids.txt")
		color.White(tokens[1])
		color.White(err.Error())
		time.Sleep(2 * time.Second)
		var wg sync.WaitGroup
		wg.Add(len(tokens))

		start := time.Now()
		for i := 0; i < len(tokens); i++ {
			time.Sleep(time.Duration(200) * time.Millisecond)

			go func(i int) {
				for j := i * (len(members) / len(tokens)); j < (i+1)*(len(members)/len(tokens)); j++ {
					status := CheckToken(tokens[i])
					if status != 200 && status != 204 && status != 429 && status != -1 {
						fmt.Println(status)

					} else {
						fmt.Println("OOKOKOKOKO")
					}
				}
			}(i)
		}
	case 2:
		color.Blue("selected 2")
	}

}

func ReadLine(filename string) ([]string, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	ex = filepath.ToSlash(ex)
	color.White(ex)
	file, err := os.Open(path.Join(path.Dir(ex) + "/input/" + filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CheckToken(auth string) int {
	url := "https://discord.com/api/v9/users/@me/affinities/guilds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1
	}
	req.Close = true
	req.Header.Set("authorization", auth)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(CommonHeaders(req))
	if err != nil {
		return -1
	}

	return resp.StatusCode
}

// Append items from slice to file
func Append(filename string, items []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, item := range items {
		if _, err = file.WriteString(item + "\n"); err != nil {
			return err
		}
	}

	return nil
}
