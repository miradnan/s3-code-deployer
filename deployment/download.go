package deployment

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fatih/color"
	"log"
	"os"
	"time"
)

// Download latest tar.gz or zip file in S3 Bucket
func DownloadPackage() bool {
	log.Println(color.CyanString(fmt.Sprintf("Downloading Package %s from S3 Bucket %s", App.RevisionFile, Config.Aws.S3Bucket)))
	// https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/s3_download_object.go

	file, err := os.Create(App.TempFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open file %s", err.Error()))
	}
	defer file.Close()

	config := &aws.Config{
		Region:      aws.String(Config.Aws.Region),
		Credentials: credentials.NewStaticCredentials(Config.Aws.AccessKey, Config.Aws.SecretKey, ""),
	}

	sess := session.New(config)

	input := s3.HeadObjectInput{
		Bucket: aws.String(Config.Aws.S3Bucket),
		Key:    aws.String(App.RevisionFile),
	}

	color.Cyan("Checking for revision version")

	result, err := s3.New(sess).HeadObject(&input)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to open file %s", err.Error()))
	}

	//log.Println(result)
	newVersion := *result.VersionId
	oldVersion := GetVersion()
	if oldVersion == newVersion {
		color.Green("No version updated")
		return false
	}

	//////////////////////////////////////////
	// Start updating the package now
	//////////////////////////////////////////

	color.MagentaString("New version found")
	color.GreenString("Downloading latest version")

	UpdateVersion(newVersion)

	totalFileSize := aws.Int64Value(result.ContentLength)
	log.Println(fmt.Sprintf("The file is %d bytes long", totalFileSize))

	downloader := s3manager.NewDownloader(sess)

	done := make(chan int64)
	go PrintDownloadPercent(done, App.TempFile, int64(totalFileSize))

	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(Config.Aws.S3Bucket),
		Key:    aws.String(App.RevisionFile),
	})

	done <- numBytes

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to download item %s, %s", App.RevisionFile, err.Error()))
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	return true
}

// PrintDownloadPercent
func PrintDownloadPercent(done chan int64, path string, total int64) {

	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			fmt.Printf("%.0f", percent)
			fmt.Println("%")
		}

		if stop {
			break
		}

		time.Sleep(time.Second * 2)
	}
}
