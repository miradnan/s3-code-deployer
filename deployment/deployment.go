package deployment

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/miradnan/codeDeployer/config"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	Config *config.Configuration
	App    *config.AppConfiguration
)

// Run the deployment
func Deploy(config *config.Configuration, app *config.AppConfiguration) {
	Config = config
	App = app

	if !CheckFileSystem() {
		log.Fatal(color.RedString(fmt.Sprintf("Invalid dir configuration in application: %s", App.Application)))
	}

	// Should only run if version is updated
	if DownloadPackage() {
		// The time it took to deploy the latest downloaded tar ball
		defer Duration()()

		// only delete tar ball / zip
		defer CleanUp()

		BeforeInstall()
		color.Green(fmt.Sprintf("Deploying in folder %s", App.DirectoryPath))

		file, err := os.Open(App.TempFile)
		if err != nil {
			log.Fatal(color.RedString(fmt.Sprintf("%s", err.Error())))
		}
		defer file.Close()
		Extract(file)

		color.Green(fmt.Sprintf("Extracting files of %s", App.DirectoryPath))

		AfterInstall()
	}
}

// CheckFileSystem
func CheckFileSystem() bool {
	color.Cyan(fmt.Sprintf("S3 Revision File: %s", App.RevisionFile))
	color.Cyan(fmt.Sprintf("App Dir: %s", App.DirectoryPath))

	// Check for the Application Directory
	if _, err := os.Stat(App.DirectoryPath); os.IsNotExist(err) {
		os.MkdirAll(App.DirectoryPath, os.ModePerm)
	}

	if _, err := os.Stat(App.TempPath); os.IsNotExist(err) {
		os.MkdirAll(App.TempPath, os.ModePerm)
	}

	return true
}

// Duration
func Duration() func() {
	start := time.Now()
	return func() {
		log.Println(color.CyanString(fmt.Sprintf("Deployment took %v", time.Since(start))))
	}
}

// BeforeInstall
func BeforeInstall() {
	ScriptPath := fmt.Sprintf("sh %s/deployment/scripts/BeforeInstall.sh", App.DirectoryPath)

	if _, err := os.Stat(ScriptPath); !os.IsNotExist(err) {
		color.Cyan(fmt.Sprintf("Running BeforeInstall: %s", ScriptPath))
		fmt.Println(exec.Command(ScriptPath).Output())
	} else {
		color.Red(fmt.Sprintf("No BeforeInstall.sh in application %s", App.Application))
	}
}

// AfterInstall
func AfterInstall() {
	ScriptPath := fmt.Sprintf("sh %s/deployment/scripts/AfterInstall.sh", App.DirectoryPath)

	if _, err := os.Stat(ScriptPath); !os.IsNotExist(err) {
		color.Cyan(fmt.Sprintf("Running AfterInstall File: %s", ScriptPath))
		fmt.Println(exec.Command(ScriptPath).Output())
	} else {
		color.Red(fmt.Sprintf("No AfterInstall.sh in application %s", App.Application))
	}
}

// Removes the previously download tar.gz or zip file
func CleanUp() {
	color.Red("Running CleanUp")
	// delete tar ball
	var err = os.Remove(App.TempFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	color.Green("CleanUp completed")
}
