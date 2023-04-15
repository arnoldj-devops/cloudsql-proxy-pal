package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Function lists the active DB connections
func listConn() {
	command := fmt.Sprintf("ps aux  | grep cloud_sql_proxy | grep -v grep | awk -F '-instances=' '{print $NF}'") //print only the cloudsql proxy process
	processlist := exec.Command("bash", "-c", command)
	output, _ := processlist.Output()
	line := strings.TrimSuffix(string(output), "\n") // removing newline
	str := strings.Split(line, "\n")                 // splitting on newline
	if str[0] == "" {                                // if no process is running
		fmt.Println("No Instance connected")
		os.Exit(1)
	}
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)
	boldGreen.Printf(strings.Join(str, "\n"))
}
