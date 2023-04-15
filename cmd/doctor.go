package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
)

func doctor() {

	// Check if gcloud sdk is available
	gcloudversioncommand := fmt.Sprintf("gcloud version | head -n 1")
	gcloudversion := exec.Command("bash", "-c", gcloudversioncommand)
	gcloudversionout, err := gcloudversion.Output()
	if err != nil {
		fmt.Println("Please reinstall gcloud sdk")
		log.Fatal(err)
	}
	fmt.Printf("gcloud version: %s", gcloudversionout)

	// Check if user is authenticated in gcloud
	gcloudusercommand := fmt.Sprintf("gcloud auth list --filter=status:ACTIVE --format='value(account)'")
	gclouduser := exec.Command("bash", "-c", gcloudusercommand)
	gclouduserout, err := gclouduser.Output()
	if err != nil {
		fmt.Println("User not authenticated\nRun: gcloud auth application-default login")
		log.Fatal(err)
	}
	fmt.Printf("Authenticated user account: %s", gclouduserout)

	// Check if cloud_sql_proxy is installed
	cloudsqlproxyversion := exec.Command("cloud_sql_proxy", "--version")
	cloudsqlproxyversionout, err := cloudsqlproxyversion.Output()
	if err != nil {
		fmt.Println("Please reinstall cloud_sql_proxy")
		log.Fatal(err)
	}
	fmt.Printf("cloud_sql_proxy version: %s", cloudsqlproxyversionout)

	//Print eveything is ok
	green := color.New(color.FgGreen)
	boldGreen := green.Add(color.Bold)
	boldGreen.Printf("Your system is ready to connect CloudSQL instances")
}
