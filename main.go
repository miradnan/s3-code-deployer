package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/miradnan/s3-code-deployer/config"
	"github.com/miradnan/s3-code-deployer/deployment"
	"io/ioutil"
	"log"
	"time"
)

var (
	Quiet   bool
	Help    bool
	Version string
	Config  *config.Configuration
	Commit  string
	BuiltAt string
	BuiltBy string
	BuiltOn string
)

func main() {

	flag.BoolVar(&Help, "h", false, "Help info")
	flag.BoolVar(&Quiet, "q", false, "Execute quietly")
	flag.StringVar(&Version, "q", "", "Get version of s3-code-deployer")

	flag.Parse()

	fmt.Print("Version info :: ")
	fmt.Printf("commit: %s ", Commit)
	fmt.Printf("built @ %s by %s on %s\n", BuiltAt, BuiltBy, BuiltOn)

	Config = config.NewConfig()

	if Quiet {
		color.Green("Running Quietly")
		log.SetOutput(ioutil.Discard)
	}

	color.Yellow("#################################################")
	color.Yellow(fmt.Sprintf("############### %s ##################", Config.Name))
	color.Yellow("#################################################")

	// Start running deployments
	for range time.NewTicker(Config.RevisionCheckDuration * time.Second).C {
		log.Println(color.CyanString(fmt.Sprintf("Executed for %d Minutes", Config.RevisionCheckDuration)))

		for i := 0; i < len(Config.Deployments); i++ {
			//go func(i int) {
			App := &Config.Deployments[i]

			if len(App.Environment) <= 0 {
				log.Fatal(color.RedString(fmt.Sprintf("Environment is required in application %s", App.Application)))
			}

			color.Yellow(fmt.Sprintf("Environment: %s", App.Environment))

			deployment.Deploy(Config, App)
			//}(i)
		}
	}
}
