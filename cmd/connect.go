package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sqladmin/v1"
)

var project string

const (
	InfoColor   = "\033[1;34m%s\033[0m"
	NoticeColor = "\033[1;36m%s\033[0m"
	GreenColor  = "\033[32m"
)

// The function gets a list of all the projects in the GCP account, then prompts the user to select one
// of the projects
func getProject() string {
	var proj_id []string
	getprojectcommand := fmt.Sprintf("gcloud projects list --format='value(project_id)'")
	getproject := exec.Command("bash", "-c", getprojectcommand)
	getprojectout, err := getproject.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		proj := strings.TrimSuffix(string(getprojectout), "\n")
		proj_id = strings.Split(proj, "\n")
	}

	prompt := promptui.Select{
		Label: "Select GCP Project",
		Items: proj_id,
		Size:  10,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
		return ""
	}

	fmt.Printf("Project ID: %q\n", result)
	promptresult := strings.Split(result, ":")
	project_id := promptresult[len(promptresult)-1]
	return project_id
}

// It uses the Google Cloud SQL Admin API to list all the instances in a project
func listInstances(project string) []string {
	var list []string
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, sqladmin.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	sqladminService, err := sqladmin.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := sqladminService.Instances.List(project)
	if err := req.Pages(ctx, func(page *sqladmin.InstancesListResponse) error {
		for _, databaseInstance := range page.Items {
			list = append(list, databaseInstance.ConnectionName)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return list
}

// It takes a list of instances, prompts the user to select one, and returns the selected instance
func getInstance() string {
	project := getProject()
	instancelist := listInstances(project)
	prompt := promptui.Select{
		Label: "Select instance to connect",
		Items: instancelist,
		Size:  10,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
		return ""
	}
	blue := color.New(color.FgBlue)
	boldBlue := blue.Add(color.Bold)
	boldBlue.Printf("You choose: %q\n", result)
	return result
}

// It starts a background process that runs the `cloud_sql_proxy` command, and then prints out the
// command to connect to the database
func connectInstance(port int) {
	var userName string
	sqlConnectionName := getInstance()
	fmt.Println("Connecting Instance")
	cmd := exec.Command("cloud_sql_proxy", "-enable_iam_login", "-instances="+sqlConnectionName+"=tcp:"+strconv.Itoa(port))
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Loading spinner
	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()

	log.Printf("Cloudsql proxy process is running in background, process_id: %d\n", cmd.Process.Pid)

	command := fmt.Sprintf("gcloud auth list --filter=status:ACTIVE --format='value(account)'")
	user := exec.Command("bash", "-c", command)
	userOut, err := user.Output()
	if err != nil {
		userName = "<username>"
	} else {
		userName = strings.TrimSuffix(string(userOut), "\n")
	}

	color.Blue("%s", "Can connect using:")
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)
	boldGreen.Printf("psql -h localhost -U %s -p %d -d postgres\n", userName, port)
}
