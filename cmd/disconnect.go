package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

// It gets the list of all the processes running and then prompts the user to select the process to be
// killed
func getPort() string {
	command := fmt.Sprintf("ps aux  | grep cloud_sql_proxy | grep -v grep | awk -F '-instances=' '{print $NF}'") //print only the cloudsql proxy process
	processlist := exec.Command("bash", "-c", command)
	output, _ := processlist.Output()
	line := strings.TrimSuffix(string(output), "\n") // removing newline
	str := strings.Split(line, "\n")                 // splitting on newline
	if str[0] == "" {                                // if no process is running
		fmt.Println("No Instance connected")
		os.Exit(1)

	}

	prompt := promptui.Select{
		Label: "Select Instance to disconnect",
		Items: str,
		Size:      10,
		Templates: &promptui.SelectTemplates{Label: "{{ . }}"},
		Searcher: func(input string, index int) bool {
			item := strings.ToLower(str[index])
			input = strings.ToLower(input)
			return strings.Contains(item, input)
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	res1 := strings.Split(result, ":")
	port := res1[len(res1)-1]
	fmt.Println(port)
	return port
}

// `lsof -i tcp:<port> | grep LISTEN | awk '{print }' | xargs kill -9`
// 
// The `lsof` command lists all open files. The `-i` flag filters the list to only include files that
// are using the specified internet address. The `tcp:<port>` flag specifies the port number. The
// `grep` command filters the list to only include files that are listening. The `awk` command prints
// the second column of the output, which is the process ID. The `xargs` command executes the `kill -9`
// command on the process ID. The `kill -9` command kills the process
func disconnectInstance() {
	port := getPort()
	command := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port) //kill the port
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Instance disconnected")
}
