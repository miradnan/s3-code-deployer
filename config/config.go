package config

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const ApplicationName = "S3CodeDeployer"

type Configuration struct {
	Name                  string             `yaml:"name"`
	Aws                   *AwsConfiguration  `yaml:"aws"`
	RevisionCheckDuration time.Duration      `yaml:"revision_check_duration"`
	Deployments           []AppConfiguration `yaml:"deployments"`
}

type AppConfiguration struct {
	Application   string `yaml:"application"`
	DirectoryPath string `yaml:"destination"`
	Environment   string `yaml:"environment"`
	RevisionFile  string `yaml:"s3_revision_file"` // latest tar ball on S3
	VersionFile   string
	TempPath      string
	TempFile      string
}

// AWS Configuration
type AwsConfiguration struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Region    string `yaml:"region"`
	S3Bucket  string `yaml:"bucket"`
}

var (
	Config *Configuration
)

func NewConfig() *Configuration {
	configFile := fmt.Sprintf("%s/config.yml", GetDeployerDir())

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatal(color.RedString("No config.yml file found"))
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(color.RedString(fmt.Sprintf("yamlFile.Get err   #%v ", err)))
	}

	Config = &Configuration{}

	err = yaml.Unmarshal(yamlFile, Config)
	if err != nil {
		panic(color.RedString(fmt.Sprintf("Unmarshal %v ", err)))
	}

	// Name
	if len(Config.Name) <= 0 {
		Config.Name = ApplicationName
	}

	// Set Values for TempPath and VersionFile
	for i := 0; i < len(Config.Deployments); i++ {
		Config.Deployments[i].TempPath = "/tmp/"
		Config.Deployments[i].TempFile = fmt.Sprintf("%s/%s", Config.Deployments[i].DirectoryPath, filepath.Base(Config.Deployments[i].RevisionFile))
		Config.Deployments[i].VersionFile = fmt.Sprintf("%s/version", Config.Deployments[i].DirectoryPath)
	}

	return Config
}

func GetDeployerDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
