package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	username, _ := reader.ReadString('\n')
	repos, err := getRepositories(username)
	if err != nil {
		fmt.Println("Error fetching repositories:", err)
		return
	}
	cloneRepositories(username, repos)
}

func getRepositories(username string) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("gh repo list %s --json name | jq .[].name -r", strings.TrimSpace(username)))

	fmt.Print(cmd)
	out, err := cmd.Output()
	fmt.Print(string(out))

	if err != nil {
		return "", err
	}
	return string(out), nil
}

func cloneRepositories(username, repos string) error {
	urls := strings.Split(repos, "\n")
	for _, url := range urls {
		url = strings.TrimSpace(url) // Remove leading/trailing whitespace, including newline characters
		url = strings.ReplaceAll(url, `"`, "")
		if url == "" {
			continue
		}
		fmt.Printf("Cloning %s\n", url)
		path := fmt.Sprintf("https://github.com/%s/%s.git", strings.TrimSpace(username), strings.TrimSpace(url))
		cmd := exec.Command("git", "clone", path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
