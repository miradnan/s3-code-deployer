package deployment

import (
	"io/ioutil"
	"log"
	"os"
)

// GetVersion
func GetVersion() string {
	dat, err := ioutil.ReadFile(App.VersionFile)
	if err != nil {
		log.Println(err.Error())

		file, err := os.OpenFile(App.VersionFile, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()
	}
	return string(dat)
}

// UpdateNewVersion
func UpdateVersion(version string) bool {
	v := []byte(version)
	err := ioutil.WriteFile(App.VersionFile, v, 0644)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
