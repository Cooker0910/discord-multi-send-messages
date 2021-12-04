package main1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/zenthangplus/goccm"
)

type jsonResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ExitSafely() {
	color.Red("\nPress ENTER to EXIT")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main1() {

	// Credits
	color.Blue("\r\n\r\n\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2584    \u2584\u2584\u2584\u2584\u2588\u2588\u2588\u2584\u2584\u2584\u2584   \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2584          \u2584\u2588\u2588\u2588\u2588\u2588\u2588\u2584   \u2584\u2588\u2588\u2588\u2588\u2588\u2588\u2584  \r\n\u2588\u2588\u2588   \u2580\u2588\u2588\u2588 \u2584\u2588\u2588\u2580\u2580\u2580\u2588\u2588\u2588\u2580\u2580\u2580\u2588\u2588\u2584 \u2588\u2588\u2588   \u2580\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2580  \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588       \u2584\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588      \u2580\u2580\u2588\u2588\u2588 \u2588\u2588\u2588\u2588\u2584  \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588   \u2584\u2588\u2588\u2588 \u2588\u2588\u2588   \u2588\u2588\u2588   \u2588\u2588\u2588 \u2588\u2588\u2588   \u2584\u2588\u2588\u2588        \u2588\u2588\u2588    \u2588\u2588\u2588 \u2588\u2588\u2588    \u2588\u2588\u2588 \r\n\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580   \u2580\u2588   \u2588\u2588\u2588   \u2588\u2580  \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580         \u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2580   \u2580\u2588\u2588\u2588\u2588\u2588\u2588\u2580  \r\n                                                                   \r\n\rDISCORD MASS DM GO")
	color.Green("\nV1.0.6 - Made by https://github.com/V4NSH4J ")
	color.Red("For educational purposes only. Read the full disclaimer and terms of use on GitHub readme file.")
	time.Sleep(2 * time.Second)
	// Check all files
	color.Green("[%v] Checking all files", time.Now().Format("15:05:04"))
	cfg, err := utilities.GetConfig()
	if err != nil {
		color.Red("Error while opening config.json: %v", err)
		ExitSafely()
		return
	}
	color.Green("[%v] Config Validated!", time.Now().Format("15:05:04"))

	msg, err := utilities.GetMessage()
	if err != nil {
		fmt.Printf("Error while opening message.json: %v", err)
		ExitSafely()
		return
	}
	color.Green("[%v] Message validated: %v\n", time.Now().Format("15:05:04"), msg)

	tkns, err := utilities.ReadLines("tokens.txt")
	if err != nil {
		color.Red("Error while opening tokens.txt: %v", err)
		ExitSafely()
		return
	}

	if len(tkns) == 0 {
		color.Red("[%v] Enter your tokens in tokens.txt")
		ExitSafely()
		return
	}
	color.Green("[%v] Tokens validated: %v tokens loaded \n", time.Now().Format("15:05:04"), len(tkns))

	members, err := utilities.ReadLines("memberids.txt")
	if err != nil {
		color.Red("Error while opening memberids.txt: %v", err)
		ExitSafely()
		return
	}

	if len(members) == 0 {
		color.Red("[%v] Enter your member ids in memberids.txt")
		ExitSafely()

		return
	}

	color.Green("[%v] Member ids validated: %v member ids loaded \n", time.Now().Format("15:05:04"), len(members))

	if cfg.Proxy != "" {
		color.Green("[%v] Now setting proxy as %v", time.Now().Format("15:05:04"), cfg.Proxy)
		os.Setenv("http_proxy", "http://"+cfg.Proxy)
		os.Setenv("https_proxy", "http://"+cfg.Proxy)
	} else {
		color.Green("[%v] Proxyless mode", time.Now().Format("15:05:04"))
	}
	// All Files validated.
	Options()
}

// Options menu
func Options() {
	color.Yellow("Leave a star on https://github.com/V4NSH4J/discord-mass-DM-GO for updates!")
	color.White("Menu:\n1) Invite Joiner\n2) Mass DM advertiser\n3) Guild Leaver\n4) Exit")
	color.White("\nEnter your choice: ")
	var choice int
	fmt.Scanln(&choice)
	if choice != 1 && choice != 2 && choice != 3 && choice != 4 && choice != 5 && choice != 6 && choice != 7 && choice != 8 && choice != 9 && choice != 0 {
		color.Red("[%v] Invalid choice", time.Now().Format("15:05:04"))

		return
	}
	switch choice {
	case 1:
		color.Cyan("Multiple Invite Mode")
		cfg, err := utilities.GetConfig()
		if err != nil {
			color.Red("Error while opening config.json: %v", err)
			ExitSafely()
			return
		}
		tokens, err := utilities.ReadLines("tokens.txt")
		if err != nil {
			color.Red("Error while opening tokens.txt: %v", err)
			ExitSafely()
			return
		}
		if len(tokens) == 0 {
			color.Red("[%v] Enter your tokens in tokens.txt", time.Now().Format("15:05:04"))
			ExitSafely()
			return
		}
		invites, err := utilities.ReadLines("invite.txt")
		if err != nil {
			color.Red("Error while opening invite.txt: %v", err)
			ExitSafely()
			return
		}
		if len(invites) == 0 {
			color.Red("[%v] Enter your invites in invite.txt", time.Now().Format("15:05:04"))
			ExitSafely()
			return
		}
		color.White("Enter delay between 2 consecutive joins by 1 token in seconds: ")
		var delay int
		fmt.Scanln(&delay)
		color.White("Enter number of Threads (0 for unlimited): ")
		var threads int
		threads = 0
		c := goccm.New(threads)
		for i := 0; i < len(tokens); i++ {
			time.Sleep(time.Duration(cfg.Offset) * time.Millisecond)
			c.Wait()
			go func(i int) {
				for j := 0; j < len(invites); j++ {
					err := utilities.Invite(invites[j], tokens[i])
					if err != nil {
						color.Red("[%v] Error while joining: %v", time.Now().Format("15:05:04"), err)
					}
					time.Sleep(time.Duration(delay) * time.Second)
				}
				c.Done()
			}(i)
		}
		c.WaitAllDone()
		color.Green("[%v] All threads finished", time.Now().Format("15:05:04"))
	case 2:
		// DM Advertiser - Blueprint
		// 1. Load all files and manage errors
		// 2. Manage the threading
		// 3. Check token {if working, continue; if not working - end instance & add IDs to failed}
		// 4. Check mutuals (if error/not mutual, continue loop & add to failed; if mutual, continue)
		// 5. Open the channel (if error, continue loop & add to failed; if no error, continue)
		// 6. Send DM (if sent, add to completed slice, print to file. If not sent; continue loop if require checking, close instance if locked)
		// 7. Truncate members with members left
		// 8. If all DMs gone, truncate tokens with tokens left
		// 9. Exit out to menu
		color.Cyan("Mass DM Advertiser/Spammer")
		color.Red("Please ensure you have used the invite joiner to join your tokens to the server and that they haven't been kicked/banned by an anti-raid bot")
		// Load files & Check for sources of error to prevent sudden crashes.
		// Also initiate variables and slices for logging and counting

		var completed []string
		var failed []string
		var dead []string
		completed, err := utilities.ReadLines("completed.txt")
		if err != nil {
			color.Red("Error while opening completed.txt: %v", err)
			ExitSafely()
			return
		}
		tokens, err := utilities.ReadLines("tokens.txt")
		if err != nil {
			color.Red("Error while opening tokens.txt: %v", err)
			ExitSafely()
			return
		}
		members, err := utilities.ReadLines("memberids.txt")
		if err != nil {
			color.Red("Error while opening members.txt: %v", err)
			ExitSafely()
			return
		}

		cfg, err := utilities.GetConfig()
		if err != nil {
			color.Red("Error while opening config.json: %v", err)
			ExitSafely()
			return
		}
		if cfg.Skip {
			members = utilities.RemoveSubset(members, completed)
		}
		msg, err := utilities.GetMessage()
		if err != nil {
			color.Red("Error while opening message.txt: %v", err)
			ExitSafely()
			return
		}
		if len(tokens) == 0 || len(members) == 0 {
			color.Red("[%v] Enter your tokens in tokens.txt and members in members.txt", time.Now().Format("15:05:04"))
			ExitSafely()
			return
		}
		if len(members) < len(tokens) {
			tokens = tokens[:len(members)]
		}
		// Threading system to prevent large-scale concurrency. Might be detectable.

		var wg sync.WaitGroup
		wg.Add(len(tokens))

		start := time.Now()
		for i := 0; i < len(tokens); i++ {
			// Offset goroutines by a few milliseconds. Made a big difference in my testing.
			time.Sleep(time.Duration(cfg.Offset) * time.Millisecond)

			go func(i int) {
				for j := i * (len(members) / len(tokens)); j < (i+1)*(len(members)/len(tokens)); j++ {
					// Check if token is still valid at start of loop. Close instance is non-functional.
					status := utilities.CheckToken(tokens[i])
					if status != 200 && status != 204 && status != 429 && status != -1 {
						color.Red("[%v] Token %v might be locked - Stopping instance and adding members to failed list. %v", time.Now().Format("15:05:04"), tokens[i], status)
						failed = append(failed, members[j:(i+1)*(len(members)/len(tokens))]...)
						dead = append(dead, tokens[i])
						err := Append("input/failed.txt", members[j:(i+1)*(len(members)/len(tokens))])
						if err != nil {
							fmt.Println(err)
						}
						if cfg.Stop {
							break
						}

					}
					var user string
					user = members[j]
					// Get user info and check for mutual servers with the victim. Continue loop if no mutual servers or error.
					// if cfg.Mutual {
					// 	info, err := directmessage.UserInfo(tokens[i], members[j])
					// 	if err != nil {
					// 		color.Red("[%v] Error while getting user info: %v", time.Now().Format("15:05:04"), err)
					// 		err = WriteLine("input/failed.txt", members[j])
					// 		if err != nil {
					// 			fmt.Println(err)
					// 		}
					// 		failed = append(failed, members[j])

					// 		continue
					// 	}
					// 	if len(info.Mutual) == 0 {
					// 		color.Red("[%v] Token %v failed to DM %v [No Mutual Server]", time.Now().Format("15:05:04"), tokens[i], info.User.Username+info.User.Discriminator)
					// 		err = WriteLine("input/failed.txt", members[j])
					// 		if err != nil {
					// 			fmt.Println(err)
					// 		}
					// 		failed = append(failed, members[j])
					// 		continue
					// 	}
					// 	user = info.User.Username + "#" + info.User.Discriminator
					// }

					// Send DM to victim. Continue loop if error.
					snowflake, err := directmessage.OpenChannel(tokens[i], members[j])
					if err != nil {
						color.Red("[%v] Error while opening DM channel: %v", time.Now().Format("15:05:04"), err)
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						failed = append(failed, members[j])
						continue
					}

					resp, err := directmessage.SendMessage(tokens[i], snowflake, &msg, members[j])
					if err != nil {
						color.Red("[%v] Error while sending message: %v", time.Now().Format("15:05:04"), err)
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						failed = append(failed, members[j])
						continue
					}
					body, err := utilities.ReadBody(*resp)
					if err != nil {
						color.Red("[%v] Error while reading body: %v", time.Now().Format("15:05:04"), err)
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						failed = append(failed, members[j])
						continue
					}
					var response jsonResponse
					errx := json.Unmarshal(body, &response)
					if errx != nil {
						color.Red("[%v] Error while unmarshalling body: %v", time.Now().Format("15:05:04"), errx)
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						failed = append(failed, members[j])
						continue
					}
					// Everything is fine, continue as usual
					if resp.StatusCode == 200 {
						err = WriteLine("input/completed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						completed = append(completed, members[j])
						color.Green("[%v] Token %v sent DM to %v [%v]", time.Now().Format("15:05:04"), tokens[i], user, len(completed))
						// Case-based error, something unusual with data enterred
					} else if resp.StatusCode == 400 {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Red("[%v] Token %v failed to DM %v Check wether the token tried to DM itself or tried sending an empty message!", time.Now().Format("15:05:04"), tokens[i], user)
						// Forbidden - Token is being rate limited
					} else if resp.StatusCode == 403 && response.Code == 40003 {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Cyan("[%v] Token %v sleeping for %v minutes!", time.Now().Format("15:05:04"), tokens[i], int(cfg.LongDelay/60))
						time.Sleep(time.Duration(cfg.LongDelay) * time.Second)
						color.Cyan("[%v] Token %v continuing!", time.Now().Format("15:05:04"), tokens[i])
						// Forbidden - DM's are closed
					} else if resp.StatusCode == 403 && response.Code == 50007 {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Red("[%v] Token %v failed to DM %v User has DMs closed or not present in server", time.Now().Format("15:05:04"), tokens[i], user)
						// Forbidden - Locked or Disabled
					} else if (resp.StatusCode == 403 && response.Code == 40002) || resp.StatusCode == 401 || resp.StatusCode == 405 {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Red("[%v] Token %v is locked or disabled. Stopping instance. %v %v", time.Now().Format("15:05:04"), tokens[i], resp.StatusCode, response.Message)
						dead = append(dead, tokens[i])
						// Stop token if locked or disabled
						if cfg.Stop {
							break
						}
						// Forbidden - Invalid token
					} else if resp.StatusCode == 403 && response.Code == 50009 {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Red("[%v] Token %v can't DM %v. It might not have bypassed community screening.", time.Now().Format("15:05:04"), tokens[i], user)
						// General case - Continue loop. If problem with instance, it will be stopped at start of loop.
					} else {
						err = WriteLine("input/failed.txt", members[j])
						if err != nil {
							fmt.Println(err)
						}
						color.Red("[%v] Token %v couldn't DM %v Error Code: %v; Status: %v; Message: %v", time.Now().Format("15:05:04"), tokens[i], user, response.Code, resp.Status, response.Message)
					}
					time.Sleep(time.Duration(cfg.Delay) * time.Second)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()

		color.Green("[%v] Threads have finished! Writing to file", time.Now().Format("15:05:04"))
		elapsed := time.Since(start)
		color.Green("[%v] DM advertisement took %v. Successfully sent DMs to %v IDs. Failed to send DMs to %v IDs. %v tokens are dis-functional & %v tokens are functioning", time.Now().Format("15:04:05"), elapsed.Seconds(), len(completed), len(failed), len(dead), len(tokens)-len(dead))
		if cfg.Remove {
			m := utilities.RemoveSubset(tokens, dead)
			err := Truncate("input/tokens.txt", m)
			if err != nil {
				fmt.Println(err)
			}
			color.Green("Updated tokens.txt")
		}
		if cfg.RemoveM {
			m := utilities.RemoveSubset(members, completed)
			err := Truncate("input/memberids.txt", m)
			if err != nil {
				fmt.Println(err)
			}
			color.Green("Updated memberids.txt")

		}

	case 3:
		// Leavs tokens from a server
		color.Cyan("Guild Leaver")
		cfg, err := utilities.GetConfig()
		if err != nil {
			color.Red("Error while opening config.json: %v", err)
			ExitSafely()
			return
		}
		tokens, err := utilities.ReadLines("tokens.txt")
		if err != nil {
			color.Red("Error while opening tokens.txt: %v", err)
			ExitSafely()
			return
		}
		if len(tokens) == 0 {
			color.Red("[%v] Enter your tokens in tokens.txt", time.Now().Format("15:05:04"))
			ExitSafely()
			return
		}
		color.White("Enter the number of threads (0 for unlimited): ")
		var threads int
		fmt.Scanln(&threads)
		if threads > len(tokens) {
			threads = len(tokens)
		}
		if threads == 0 {
			threads = len(tokens)
		}
		color.White("Enter delay between leaves: ")
		var delay int
		fmt.Scanln(&delay)
		color.White("Enter serverid: ")
		var serverid string
		fmt.Scanln(&serverid)
		c := goccm.New(threads)
		for i := 0; i < len(tokens); i++ {
			time.Sleep(time.Duration(cfg.Offset) * time.Millisecond)
			c.Wait()
			go func(i int) {
				p := utilities.Leave(serverid, tokens[i])
				if p == 0 {
					color.Red("[%v] Error while leaving", time.Now().Format("15:05:04"))
				}
				if p == 200 || p == 204 {
					color.Green("[%v] Left server", time.Now().Format("15:05:04"))
				} else {
					color.Red("[%v] Error while leaving", time.Now().Format("15:05:04"))
				}
				time.Sleep(time.Duration(delay) * time.Second)
				c.Done()
			}(i)
		}
		c.WaitAllDone()
		color.Green("[%v] All threads finished", time.Now().Format("15:05:04"))

	case 4:
		// Exit without error
		os.Exit(0)

	}
	time.Sleep(1 * time.Second)
	Options()

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

// Truncate items from slice to file
func Truncate(filename string, items []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)
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

// Write line to file
func WriteLine(filename string, line string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}
